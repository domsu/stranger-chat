module "config" {
  source="./config"
}

provider "google" {
  credentials = "${file("./config/google_cloud_credentials.json")}"
  project = "${module.config.google_project_id}"
  region = "us-west1"
}

resource "google_compute_instance" "stranger-services" {
  name = "stranger-services"
  machine_type = "g1-small"
  zone = "us-west1-a"

  tags = [
    "service"
  ]

  boot_disk {
    initialize_params {
      image = "coreos-cloud/coreos-stable"
    }
  }

  metadata_startup_script = "sudo apt-get update;"

  network_interface {
    network = "default"

    access_config {
      nat_ip = "${google_compute_address.static.address}"
    }
  }

  metadata {
    sshKeys = "service:${file("~/.ssh/id_rsa.pub")}"
  }
}

resource "google_compute_address" "static" {
  name = "ipv4-address"
}

resource "google_dns_managed_zone" "myzone" {
  name        = "myzone"
  dns_name    = "${module.config.domain}."
}

resource "google_dns_record_set" "strangerchat-net-set" {
  name = "strangerchat.${google_dns_managed_zone.myzone.dns_name}"
  managed_zone = "${google_dns_managed_zone.myzone.name}"
  type = "A"
  ttl  = 300

  rrdatas = ["${google_compute_address.static.address}"]
}

resource "google_dns_record_set" "www-strangerchat-net-set" {
  name = "www.strangerchat.${google_dns_managed_zone.myzone.dns_name}"
  managed_zone = "${google_dns_managed_zone.myzone.name}"
  type = "A"
  ttl  = 300

  rrdatas = ["${google_compute_address.static.address}"]
}


resource "google_compute_firewall" "default" {
  name = "service-firewall"
  network = "default"

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports = [
      "443"
    ]
  }

  allow {
    protocol = "tcp"
    ports = [
      "80"
    ]
  }

  source_ranges = [
    "0.0.0.0/0"
  ]
  target_tags = [
    "service"
  ]
}

output "service-ip" {
  value = "${google_compute_address.static.address}"
}