import * as React from 'react'
import { ChakraProvider, theme } from '@chakra-ui/react'
import { Route, Routes } from 'react-router-dom'
import SignUp from './components/pages/SignUp'
import LogIn from './components/pages/LogIn'
import { AppHeader } from './components/organisms/AppHeader'
import Home from './components/pages/Home'

export const App = () => (
    <ChakraProvider theme={theme}>
        <AppHeader />
        <Routes>
            <Route path="sign-up" element={<SignUp />} />
            <Route path="login" element={<LogIn />} />
            <Route path="/" element={<Home />} />
        </Routes>
    </ChakraProvider>
)
