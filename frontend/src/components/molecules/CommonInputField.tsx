import React from 'react'
import { Box, FormLabel, Input, Text } from '@chakra-ui/react'
import { Nullable } from '../../types/utils'

type CommonInputFieldPropsType = {
    inputType: 'email' | 'password' | 'text'
    label: string
    placeholder: string
    onInputAction: React.ChangeEventHandler<HTMLInputElement>
    mt: number
    isRequired: boolean
    name: string
    value: string
    error: Nullable<string>
}

const CommonInputField = ({
    inputType,
    label,
    placeholder,
    onInputAction,
    mt,
    isRequired,
    name,
    value,
    error
}: CommonInputFieldPropsType) => {
    return (
        <Box mt={mt}>
            <FormLabel textTransform="capitalize" color="gray.600" fontWeight={600}>
                {label}
                {isRequired ? (
                    <Text as="span" textColor="red.500">
                        &nbsp;*
                    </Text>
                ) : null}
            </FormLabel>
            <Input
                type={inputType}
                placeholder={placeholder}
                onChange={onInputAction}
                borderRadius={4}
                required={isRequired}
                color="gray.500"
                name={name}
                value={value}
            />
            {error ? (
                <Text fontSize="1rem" color="red.700" fontWeight={600} lineHeight={1.2} mt={2} mb={3}>
                    {error}
                </Text>
            ) : null}
        </Box>
    )
}

export default CommonInputField
