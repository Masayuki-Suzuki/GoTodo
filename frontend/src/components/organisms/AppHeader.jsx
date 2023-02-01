import * as React from 'react'
import { Flex } from '@chakra-ui/react'
import { AppTitle } from '../atoms/AppTitle'

export const AppHeader = () => (
    <Flex w="100%" bg="#19448e" p={4} boxShadow="md" position="absolute" top={0} left={0}>
        <AppTitle />
    </Flex>
)
