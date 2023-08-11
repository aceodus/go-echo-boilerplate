package subscribers

import (
	"github.com/aceodus/go-echo-boilerplate/pkg/hook"
)

type UserSubscriber struct {
}

func (s UserSubscriber) Created(payload hook.EventPayload) {

}

func (s UserSubscriber) Deleted(payload hook.EventPayload) {

}

func (s UserSubscriber) Updated(payload hook.EventPayload) {

}
