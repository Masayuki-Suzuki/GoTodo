import React, { useEffect } from 'react'
import { Flex, Heading } from '@chakra-ui/react'
import { useNavigate } from 'react-router-dom'

const Home = () => {
    const navigate = useNavigate()

    useEffect(() => {
        const token = sessionStorage.getItem('token')
        if (!token) {
            navigate('/login')
        }
    }, [])

    return (
        <Flex w="100%" h="100vh" display="flex" align="center" justify="center" pt="56px">
            <Heading as="h1" size="2xl">
                Index Page.
            </Heading>
        </Flex>
    )
}

export default Home
