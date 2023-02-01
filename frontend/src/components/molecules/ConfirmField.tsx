import React from 'react'
import { Box, FormLabel, Text } from '@chakra-ui/react'

type ConfirmFieldPropsType = {
    label: string
    mt: number
    email: string
}

const ConfirmField = ({ mt, label, email }: ConfirmFieldPropsType) => {
    return (
        <Box mt={mt}>
            <FormLabel textTransform="capitalize" color="gray.600" fontWeight={600}>
                {label}:
            </FormLabel>
            <Text fontSize="1.25rem" fontWeight={400} color="gray.500" minH={30}>
                {email}
            </Text>
        </Box>
    )
}

export default ConfirmField
