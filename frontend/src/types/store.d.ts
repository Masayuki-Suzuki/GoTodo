import { User } from '../reducks/users/types'
import { FetchStatus } from '../reducks/fetchStatus/types'
import { ModalState } from '../reducks/modals/types'
import { ToDos } from '../reducks/ToDos/types'

export type RootState = {
    user: User
    fetchStatus: FetchStatus
    modalState: ModalState
    todos: ToDos
}
