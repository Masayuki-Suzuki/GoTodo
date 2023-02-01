import * as React from 'react'
import { ChakraProvider, theme } from '@chakra-ui/react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import SignUp from './components/pages/SignUp'
import { AppHeader } from './components/organisms/AppHeader'

export const App = () => (
    <ChakraProvider theme={theme}>
        <AppHeader />
        <Routes>
            <Route path="sign-up" element={<SignUp />} />
        </Routes>
    </ChakraProvider>
)
