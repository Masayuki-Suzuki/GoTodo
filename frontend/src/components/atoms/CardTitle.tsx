import React from 'react'
import { Heading } from '@chakra-ui/react'

type PropsType = {
    title: string
}

export const CardTitle = ({ title }: PropsType) => (
    <Heading as="h3" fontSize="2xl" color="gray.600" fontWeight={500} textAlign="center" w="100%" px={6}>
        {title}
    </Heading>
)
