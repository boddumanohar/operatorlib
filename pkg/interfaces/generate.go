//go:generate mockgen -destination=./mocks/object.go -package=mocks github.com/ankitrgadiya/operatorlib/pkg/interfaces Object
//go:generate mockgen -destination=./mocks/reconcile.go -package=mocks github.com/ankitrgadiya/operatorlib/pkg/interfaces Reconcile

package interfaces
