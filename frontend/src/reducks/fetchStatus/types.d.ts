import { Nullable } from '~/types/utils'

export type FetchStatus = {
    isLoading: boolean
    error: Nullable<string>
}
