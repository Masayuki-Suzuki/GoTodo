import { createSelector } from 'reselect'
import { RootState } from '~/types/store'

const fetchStatusSelector = (state: RootState) => state.fetchStatus

export const getFetchStatus = createSelector([fetchStatusSelector], state => state.isLoading)

export default {
    getFetchStatus
}
