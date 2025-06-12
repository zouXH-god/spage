"use client";

import { useState, useEffect } from "react";
import { login, getCaptchaConfig } from "@/api/user.api";
import AIOCaptchaWidget from "@/components/captcha/AIOCaptcha";
import { CaptchaProps, CaptchaProvider } from "@/types/captcha";
import { t } from "i18next";
import { useDevice } from "@/contexts/DeviceContext";

export default function LoginPage() {
  const { isMobile } = useDevice();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [captchaToken, setCaptchaToken] = useState("");
  const [captchaProps, setCaptchaProps] = useState<CaptchaProps | null>(null);

  useEffect(() => {
    getCaptchaConfig()
      .then((response) => {
        const config = response.data;
        setCaptchaProps({
          provider: config.provider as CaptchaProvider,
          siteKey: config.siteKey || "",
          url: config.url || "",
          onSuccess: (token: string) => setCaptchaToken(token),
          onError: () => setError("验证码验证失败，请重试"),
        });
      })
      .catch(() => setError("获取验证码配置失败，请稍后再试"));
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    login({ username, password, captchaToken })
      .then((response) => {
        console.log("登录成功:", response.data);
        // 处理登录成功逻辑
      })
      .catch(() => setError("登录失败，请检查信息后重试"));
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 dark:bg-gray-900">
      <form
        onSubmit={handleSubmit}
        className={
          isMobile
            ? "w-full min-h-screen flex flex-col justify-center px-4 bg-white dark:bg-gray-800"
            : "w-[400px] h-[420px] bg-white dark:bg-gray-800 rounded-2xl shadow-2xl p-10 flex flex-col justify-center"
        }
        style={isMobile ? {} : { minWidth: 320, minHeight: 380 }}
      >
        <h2 className="text-3xl font-bold text-center text-gray-900 dark:text-white mb-6 tracking-tight">
          {t("login.login")}
        </h2>
        <div className="space-y-4 mb-4">
          <input
            type="text"
            placeholder={t("login.username")}
            value={username}
            onChange={e => setUsername(e.target.value)}
            className="w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-700 bg-gray-100 dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <input
            type="password"
            placeholder={t("login.password")}
            value={password}
            onChange={e => setPassword(e.target.value)}
            className="w-full px-4 py-2 rounded-lg border border-gray-300 dark:border-gray-700 bg-gray-100 dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div className="flex justify-center my-4">
          <AIOCaptchaWidget
            {...(captchaProps || {
              provider: CaptchaProvider.DISABLE,
              siteKey: "",
              url: "",
              onSuccess: () => { },
              onError: () => { },
            })}
          />
        </div>
        {error && (
          <div className="text-red-500 text-sm text-center mb-2">{error}</div>
        )}
        <button
          type="submit"
          disabled={!captchaToken}
          className={
            "w-full py-2 rounded-lg font-semibold text-lg transition-colors mt-2 " +
            (captchaToken
              ? "bg-blue-600 hover:bg-blue-700 text-white cursor-pointer"
              : "bg-blue-200 text-white cursor-not-allowed")
          }
        >
          {t("login.login")}
        </button>
      </form>
    </div>
  );
}