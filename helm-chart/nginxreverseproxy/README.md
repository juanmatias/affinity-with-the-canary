# Affinity with the canary - HELM Chart

I must write this page!

## Subchart

You can use this chart as a subchart for your own chart.

In your *requirements.yaml*:

    dependencies:
      - name: affinitwiththecanary
        version: 0.1.1
        repository: https://raw.githubusercontent.com/juanmatias/affinity-with-the-canary/master

In your *values.yaml*:

    env:
      # baseSvcUrl: ""
      # these two should be {{ .Values.service.name }} and {{ .Values.service.name }}-canary
      stableSvcName: svcname
      canarySvcName: svcname-canary
      stableSvcPort: "8080"
      canarySvcPort: "8080"
      # from here on only valid if canary.ingress.enabled=false
      canaryWeight: "0.5"
      cookieMaxAge: "172800"
      cookieName: "cookiemonster"
      cookieEnabled: "true"
    service:
      name: nginxcan-service
      type: ClusterIP
      port: 80
    autoscaling:
      hpa:
        enabled: false
