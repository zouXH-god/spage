import { useEffect } from "react";

import { CaptchaProps } from "@/types/captcha";

export default function NoCaptchaWidget(props: CaptchaProps) {
  useEffect(() => {
    props.onSuccess("no-captcha");
  }, [props, props.onSuccess]);
  return null;
}
