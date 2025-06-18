import { CaptchaProps } from "@/types/captcha";
import HCaptcha from "@hcaptcha/react-hcaptcha";

export default function HCaptchaWidget(props: CaptchaProps) {
  return (
    <HCaptcha
      sitekey={props.siteKey}
      onVerify={props.onSuccess}
      onError={props.onError}
      onExpire={() => props.onError?.("Captcha expired")}
    />
  );
}
