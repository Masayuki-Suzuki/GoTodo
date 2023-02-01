import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { FetchStatus } from './types'

const initialState: FetchStatus = {
    isLoading: false
}

export const fetchStatusSlice = createSlice({
    name: 'fetchStatus',
    initialState,
    reducers: {
        loadingStatusAction: (state: FetchStatus, action: PayloadAction<FetchStatus>) => {
            const updatedData = { ...state, ...action.payload }
            return updatedData
        }
    }
})

export const { loadingStatusAction } = fetchStatusSlice.actions
export const fetchStatus = fetchStatusSlice.reducer
