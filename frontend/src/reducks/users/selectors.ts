import { createSelector } from 'reselect'
import { RootState } from '../../types/store'

const userSelector = (state: RootState) => state.user

export const getUserId = createSelector([userSelector], state => state.id)

export default {
    getUserId
}
