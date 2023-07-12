# GPU Exporter

The Nvidia Graphics Card Metrics Exporter project aims to provide a seamless solution for extracting and monitoring metrics from Nvidia graphics cards. By utilizing Prometheus as the monitoring system and Consul for service discovery, this exporter enables efficient collection and consumption of valuable metrics from Nvidia GPUs.

The primary objective of this project is to simplify the process of monitoring and analyzing the performance of Nvidia graphics cards within a distributed system. By exporting metrics through the Prometheus exposition format, the exporter enables compatibility with the rich ecosystem of Prometheus-based tools and integrations. This allows users to leverage the extensive monitoring capabilities provided by Prometheus, such as visualization, alerting, and long-term data retention.

To facilitate seamless integration into diverse environments, the Nvidia Graphics Card Metrics Exporter comes pre-integrated with Consul, a popular service discovery and configuration management tool. This integration enables automatic discovery of the exporter instances, ensuring that the metrics from all Nvidia graphics cards within the system can be efficiently collected and monitored.

By combining the power of Prometheus and Consul, users can effortlessly gather and visualize metrics from multiple Nvidia graphics cards in a centralized manner. The exporter's integration with Consul simplifies the setup process by automatically detecting new client instances as they join or leave the system, ensuring the continuous availability of metrics.

## Key Features:

- Export metrics from Nvidia graphics cards in the Prometheus exposition format.
- Seamless integration with Prometheus for advanced monitoring capabilities.
- Automatic service discovery of exporter instances using Consul.
- Scalable and efficient collection of metrics from multiple Nvidia GPUs.
- Simplified setup and configuration through Consul's service discovery capabilities.

## Coming

- MTLS
- Authentication
- CPU, RAM metrics