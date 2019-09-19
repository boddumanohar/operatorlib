//go:generate mockgen -destination=./mocks/object.go -package=mocks argp.in/go/operatorlib/pkg/interfaces Object
//go:generate mockgen -destination=./mocks/reconcile.go -package=mocks argp.in/go/operatorlib/pkg/interfaces Reconcile

package interfaces
