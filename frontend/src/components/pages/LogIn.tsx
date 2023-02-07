import React, { useEffect } from 'react'
import { Flex } from '@chakra-ui/react'
import { useNavigate } from 'react-router-dom'
import LogInCard from '../organisms/LogInCard'

const LogIn = () => {
    const navigate = useNavigate()

    useEffect(() => {
        const token = sessionStorage.getItem('token')
        if (token) {
            navigate('/')
        }
    }, [])

    return (
        <Flex w="100%" h="100vh" display="flex" align="center" justify="center" pt="56px">
            <LogInCard />
        </Flex>
    )
}

export default LogIn
