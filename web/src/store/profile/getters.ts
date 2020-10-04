import { GetterTree } from 'vuex'
import { ProfileState } from './ProfileState'
import { RootState } from './RootState'

export const getters: GetterTree<ProfileState, RootState> = {
  fullName (state): string {
    const { user } = state
    const firstName = (user && user.idToken && user.idToken.given_name) || ''
    const lastName = (user && user.idToken && user.idToken.family_name) || ''
    return `${firstName} ${lastName}`
  }
}
