import { User } from '../reducks/users/types'
import { FetchStatus } from '../reducks/fetchStatus/types'

export type RootState = {
    user: User
    fetchStatus: FetchStatus
}
