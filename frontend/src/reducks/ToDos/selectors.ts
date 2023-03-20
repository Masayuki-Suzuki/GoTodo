import { createSelector } from 'reselect'
import { RootState } from '../../types/store'

const todoSelector = (state: RootState) => state.todos
export const getToDos = createSelector([todoSelector], state => state)

export default {
    getToDos
}
