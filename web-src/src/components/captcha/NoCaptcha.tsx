import { CaptchaProps } from '@/types/captcha'
import { useEffect } from 'react'

export default function NoCaptchaWidget(props: CaptchaProps) {
    useEffect(() => {
        props.onSuccess("no-captcha")
    }, [props.onSuccess])
    return null
}