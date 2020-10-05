import { GetterTree } from 'vuex'
import {
  Options,
  url
} from 'gravatar'
import { ProfileState } from './ProfileState'
import { RootState } from './RootState'

export const getters: GetterTree<ProfileState, RootState> = {
  fullName (state): string {
    const { user } = state
    const firstName = (user && user.idToken && user.idToken.given_name) || ''
    const lastName = (user && user.idToken && user.idToken.family_name) || ''
    return `${firstName} ${lastName}`
  },

  gravatarImage (state): string {
    const { user } = state
    const email = (user && user.idToken && user.idToken.emails[0]) || ''
    return url(email)
  }
}
