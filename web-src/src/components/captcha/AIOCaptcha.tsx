import { CaptchaProvider, type CaptchaProps } from "@/types/captcha";
import HCaptchaWidget from "./HCaptcha";
import ReCaptchaWidget from "./ReCaptcha";
import TurnstileWidget from "./Turnstile";


export default function AIOCaptchaWidget(props: CaptchaProps) {
    switch (props.provider) {
        case CaptchaProvider.HCAPTCHA:
            return <HCaptchaWidget {...props} />;
        case CaptchaProvider.RECAPTCHA:
            return <ReCaptchaWidget {...props} />;
        case CaptchaProvider.TURNSTILE:
            return <TurnstileWidget {...props} />;
        case CaptchaProvider.DISABLE:
            return null;
        default:
            throw new Error(`Unsupported captcha provider: ${props.provider}`);
    }
}