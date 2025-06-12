import {
  GoogleReCaptchaProvider,
  GoogleReCaptcha
} from 'react-google-recaptcha-v3';
import { CaptchaProps } from '@/types/captcha';

export default function ReCaptchaWidget(props: CaptchaProps) {
  return (
    <GoogleReCaptchaProvider
      reCaptchaKey={props.siteKey}
      useEnterprise={false}
    >
      <GoogleReCaptcha
        action="submit"
        onVerify={props.onSuccess}
      />
    </GoogleReCaptchaProvider>
  );
}