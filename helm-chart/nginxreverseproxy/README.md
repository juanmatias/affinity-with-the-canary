# Affinity with the canary - HELM Chart

I must write this page!

## Subchart

You can use this chart as a subchart for your own chart.

In your *requirements.yaml*:

    dependencies:
      - name: nginxreverseproxy
        version: 0.1.4
        repository: https://raw.githubusercontent.com/juanmatias/affinity-with-the-canary/master

In your project's *values.yaml*:

    nginxreverseproxy:
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
        # Enable availability zones even pod distribution
        azs:
          enabled: false
          values:
            - us-east-1a
            - us-east-1b

## Multi AZ deployment

In versions 0.1.4+ there is the feature to deploy multiaz, this means multi availability zone. 

When you work with providers such AWS, you can allocate nodes on more than one AZ. This ensures that you have pods evenly deployed in each AZ.

For more info see [here]( https://juanmatiasdelacamara.wordpress.com/2019/05/29/multi-az-deployments/).
