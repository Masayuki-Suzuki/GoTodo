import { User } from '../reducks/users/types'
import { FetchStatus } from '../reducks/fetchStatus/types'
import { ModalState } from '../reducks/modals/types'

export type RootState = {
    user: User
    fetchStatus: FetchStatus
    modalState: ModalState
}
