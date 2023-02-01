import React from 'react'
import { Flex } from '@chakra-ui/react'
import SignUpCard from '../organisms/SignUpCard'

const SignUpPage = () => {
    return (
        <Flex w="100%" h="100vh" display="flex" align="center" justify="center" pt="56px">
            <SignUpCard />
        </Flex>
    )
}

export default SignUpPage
