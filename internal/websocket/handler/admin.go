package handler

import (
	"demodesk/neko/internal/types"
	"demodesk/neko/internal/types/event"
	"demodesk/neko/internal/types/message"
)

func (h *MessageHandlerCtx) adminLock(session types.Session) error {
	if !session.Admin() {
		h.logger.Debug().Msg("user not admin")
		return nil
	}

	if h.locked {
		h.logger.Debug().Msg("server already locked...")
		return nil
	}

	h.locked = true

	h.sessions.Broadcast(
		message.Admin{
			Event: event.ADMIN_LOCK,
			ID:    session.ID(),
		}, nil)

	return nil
}

func (h *MessageHandlerCtx) adminUnlock(session types.Session) error {
	if !session.Admin() {
		h.logger.Debug().Msg("user not admin")
		return nil
	}

	if !h.locked {
		h.logger.Debug().Msg("server not locked...")
		return nil
	}

	h.locked = false

	h.sessions.Broadcast(
		message.Admin{
			Event: event.ADMIN_UNLOCK,
			ID:    session.ID(),
		}, nil)

	return nil
}

func (h *MessageHandlerCtx) adminControl(session types.Session) error {
	if !session.Admin() {
		h.logger.Debug().Msg("user not admin")
		return nil
	}

	host := h.sessions.GetHost()
	h.sessions.SetHost(session)

	if host != nil {
		h.sessions.Broadcast(
			message.AdminTarget{
				Event:  event.ADMIN_CONTROL,
				ID:     session.ID(),
				Target: host.ID(),
			}, nil)
	} else {
		h.sessions.Broadcast(
			message.Admin{
				Event: event.ADMIN_CONTROL,
				ID:    session.ID(),
			}, nil)
	}

	return nil
}

func (h *MessageHandlerCtx) adminRelease(session types.Session) error {
	if !session.Admin() {
		h.logger.Debug().Msg("user not admin")
		return nil
	}

	host := h.sessions.GetHost()
	h.sessions.ClearHost()

	if host != nil {
		h.sessions.Broadcast(
			message.AdminTarget{
				Event:  event.ADMIN_RELEASE,
				ID:     session.ID(),
				Target: host.ID(),
			}, nil)
	} else {
		h.sessions.Broadcast(
			message.Admin{
				Event: event.ADMIN_RELEASE,
				ID:    session.ID(),
			}, nil)
	}

	return nil
}

func (h *MessageHandlerCtx) adminGive(session types.Session, payload *message.Admin) error {
	if !session.Admin() {
		h.logger.Debug().Msg("user not admin")
		return nil
	}

	target, ok := h.sessions.Get(payload.ID)
	if !ok {
		h.logger.Debug().Str("id", payload.ID).Msg("can't find target session")
		return nil
	}

	h.sessions.SetHost(target)

	h.sessions.Broadcast(
		message.AdminTarget{
			Event:  event.CONTROL_GIVE,
			ID:     session.ID(),
			Target: target.ID(),
		}, nil)

	return nil
}

func (h *MessageHandlerCtx) adminKick(session types.Session, payload *message.Admin) error {
	if !session.Admin() {
		h.logger.Debug().Msg("user not admin")
		return nil
	}

	target, ok := h.sessions.Get(payload.ID)
	if !ok {
		h.logger.Debug().Str("id", payload.ID).Msg("can't find target session")
		return nil
	}

	if target.Admin() {
		h.logger.Debug().Msg("target is an admin, baling")
		return nil
	}

	if err := target.Disconnect("kicked"); err != nil {
		return err
	}

	h.sessions.Broadcast(
		message.AdminTarget{
			Event:  event.ADMIN_KICK,
			Target: target.ID(),
			ID:     session.ID(),
		}, []string{payload.ID})

	return nil
}