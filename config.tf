provider "docker" {
  host = "tcp://localhost:2375"
}
resource "docker_image" "elastic" {
  name = "docker.elastic.co/elasticsearch/elasticsearch:7.2.0"
  
}
