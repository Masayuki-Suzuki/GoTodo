import { createSelector } from 'reselect'
import { RootState } from '../../types/store'

const modalStateSelector = (state: RootState) => state.modalState

export const getModalStatus = createSelector([modalStateSelector], state => state.isOpen)

export default { getModalStatus }
