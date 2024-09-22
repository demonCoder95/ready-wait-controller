# Ready Wait Controller
A controller that waits for a given set of resources to be ready and prints the output.

It targets `apps/v1` resources. The objective of the controller is to be able to monitor
if all the provided `apps/v1` types like `StatefulSet`, `DaemonSet` and `Deployment` have
all their replicas up-to-date.

The use case is to be able to wait for an update to complete before proceeding with a provisioning
tools like the [`cluster-lifecycle-manager`](https://github.com/zalando-incubator/cluster-lifecycle-manager) project of Zalando.

TODO: The controller can either be imported as a package in your Go app or run as a CLI program.
