import { configureStore } from '@reduxjs/toolkit'
import { user } from '../users/slices'
import { fetchStatus } from '../fetchStatus/slices'

export const store = configureStore({
    reducer: {
        user,
        fetchStatus
    }
})
