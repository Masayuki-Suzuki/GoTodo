import { configureStore } from '@reduxjs/toolkit'
import { user } from '../users/slices'
import { fetchStatus } from '../fetchStatus/slices'
import { modalState } from '../modals/slices'

export const store = configureStore({
    reducer: {
        user,
        fetchStatus,
        modalState
    }
})
