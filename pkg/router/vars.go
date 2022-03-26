package router

var ClusterRole string = `apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: openshift-ingress-router
rules:
- apiGroups:
  - ""
  resources:
  - namespaces
  - services
  - endpoints
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - route.openshift.io
  resources:
  - routes
  verbs:
  - list
  - watch
- apiGroups:
  - route.openshift.io
  resources:
  - routes/status
  verbs:
  - get
  - patch
  - update
`

var ClusterRoleBindingSA string = `apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: openshift-ingress-router
roleRef:
  apiGroup: ""
  kind: ClusterRole
  name: openshift-ingress-router
subjects:
- kind: ServiceAccount
  namespace: openshift-ingress
  name: ingress-router
`

var ClusterRoleBindingDelegator string = `apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: openshift-ingress-router-auth-delegator
roleRef:
  apiGroup: ""
  kind: ClusterRole
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  namespace: openshift-ingress
  name: ingress-router
`

var Namespace string = `apiVersion: v1
kind: Namespace
metadata:
  name: openshift-ingress
`

var ServiceAccount string = `apiVersion: v1
kind: ServiceAccount
metadata:
  name: ingress-router
  namespace: openshift-ingress
`

var RouterCrd string = `apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: routes.route.openshift.io
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: route.openshift.io
  # list of versions supported by this CustomResourceDefinition
  versions:
    - name: v1
      # Each version can be enabled/disabled by Served flag.
      served: true
      # One and only one version must be marked as the storage version.
      storage: true
      schema:
        openAPIV3Schema:
          type: object
          x-kubernetes-preserve-unknown-fields: true
      additionalPrinterColumns:
        - name: Host
          type: string
          jsonPath: .status.ingress[0].host
        - name: Admitted
          type: string
          jsonPath: .status.ingress[0].conditions[?(@.type=="Admitted")].status
        - name: Service
          type: string
          jsonPath: .spec.to.name
        - name: TLS
          type: string
          jsonPath: .spec.tls.type
      subresources:
        # enable spec/status
        status: {}
  # either Namespaced or Cluster
  scope: Namespaced
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: routes
    # singular name to be used as an alias on the CLI and for display
    singular: route
    # kind is normally the CamelCased singular type. Your resource manifests use this.
    kind: Route
`

var RouterDeploy string = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: ingress-router
  namespace: openshift-ingress
  labels:
    k8s-app: ingress-router
spec:
  selector:
    matchLabels:
      k8s-app: ingress-router
  template:
    metadata:
      labels:
        k8s-app: ingress-router
    spec:
      serviceAccountName: ingress-router
      containers:
      - env:
        - name: ROUTER_LISTEN_ADDR
          value: 0.0.0.0:1936
        - name: ROUTER_METRICS_TYPE
          value: haproxy
        - name: ROUTER_SERVICE_HTTPS_PORT
          value: "443"
        - name: ROUTER_SERVICE_HTTP_PORT
          value: "80"
        - name: ROUTER_THREADS
          value: "4"
        - name: ROUTER_SUBDOMAIN
          value: "${name}-${namespace}.apps.127.0.0.1.nip.io"
        - name: ROUTER_ALLOW_WILDCARD_ROUTES
          value: "true"
        image: openshift/origin-haproxy-router:v4.0.0
        livenessProbe:
          httpGet:
            host: localhost
            path: /healthz
            port: 1936
          initialDelaySeconds: 10
        name: router
        ports:
        - containerPort: 80
        - containerPort: 443
        - containerPort: 1936
          name: stats
          protocol: TCP
        readinessProbe:
          httpGet:
            host: localhost
            path: healthz/ready
            port: 1936
          initialDelaySeconds: 10
        resources:
          requests:
            cpu: 100m
            memory: 256Mi
      hostNetwork: true
`
