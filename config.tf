provider "docker" {
  host = "tcp://localhost:2375"
}
resource "docker_image" "elastic_image" {
  name = "docker.elastic.co/elasticsearch/elasticsearch:7.2.0"
}
resource "docker_image" "logstash_image" {
  name = "docker.elastic.co/logstash/logstash:7.2.0"
}
resource "docker_image" "kibana_image" {
  name = " docker.elastic.co/kibana/kibana:7.2.0"
}
resource "docker_container" "elasticsearch" {
  name = "elasticsearch"
  image = "${docker_image.elastic_image}"
  ports {
    internal = 9200
    external = 9200
  }
    ports {
    internal = 9300
    external = 9300
  }
  env = ["discovery.type=single-node"]
}

resource "docker_container" "elastic" {
  name = "elastic"
  image = "${docker_image.logstash_image}"
  ports {
    internal = 5044
    external = 5044
  }
    ports {
    internal = 9600
    external = 9600
  }
  volumes {
    container_path = "/usr/share/logstash/pipeline/"
    host_path = "./logstash"
    read_only = true
  }
}

resource "docker_Container" "kibana" {
  name = "kibana"
  image = "${docker_image.kibana_image}"
  ports {
    internal = 5601
    external = 5601
  }
  env = ["SERVER_NAME=kibane", "ELASTICSEARCH_HOSTS=http://elasticsearch:9200"]
}
