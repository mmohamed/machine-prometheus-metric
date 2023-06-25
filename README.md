
# Machine Prometheus Metrics (exporter)

## Deploy
```bash
# deploy the exporter
kubectl apply -f sample-daemonset.yaml
```

## Metrics

| Metric name | Metric type | Labels |
|-------------|-------------|-------------|
|cpu_heat_celsius|Gauge|pod=\<pod-name\> <br/> node=\<node-name\>| 


## Build
```bash
docker build --tag machine-prometheus-metric:local . -f Dockerfile
# For multi plateform 
# docker buildx build --push --platform linux/arm/v7,linux/arm64,linux/amd64 --tag medinvention/machine-prometheus-metric:0.0.1 . -f Dockerfile
```

### References
- https://github.com/fstab/grok_exporter