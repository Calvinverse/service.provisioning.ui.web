import { EventBus } from 'estacion/lib/event-bus'

export interface EventBusServiceDefinition {
  getBus(): EventBus;
}

// This should be a singleton because internally Msal.UserAgentApplication keeps state
// which isn't linked to vuex or anything
export class EventBusService implements EventBusServiceDefinition {
  private static instance: EventBusService

  static getInstance (): EventBusService {
    if (!EventBusService.instance) {
      EventBusService.instance = new EventBusService()
    }

    return EventBusService.instance
  }

  private bus: EventBus

  private constructor () {
    this.bus = new EventBus()
  }

  getBus (): EventBus {
    return this.bus
  }
}
