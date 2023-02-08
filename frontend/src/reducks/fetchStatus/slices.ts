import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { FetchStatus } from './types'

const initialState: FetchStatus = {
    isLoading: false,
    error: null
}

export const fetchStatusSlice = createSlice({
    name: 'fetchStatus',
    initialState,
    reducers: {
        loadingStatusAction: (state: FetchStatus, action: PayloadAction<Omit<FetchStatus, 'error'>>) => {
            const updatedData = { ...state, ...action.payload }
            return updatedData
        },
        setFetchError: (state: FetchStatus, action: PayloadAction<Omit<FetchStatus, 'isLoading'>>) => {
            const updatedData = { ...state, ...action.payload }
            return updatedData
        },
        resetFetchError: (state: FetchStatus) => {
            const updateData = { ...state, error: null }
            return updateData
        }
    }
})

export const { loadingStatusAction, setFetchError, resetFetchError } = fetchStatusSlice.actions
export const fetchStatus = fetchStatusSlice.reducer
