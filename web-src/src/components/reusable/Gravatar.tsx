"use client";

import React from "react";
import Image from "next/image";
import Gravatar from "react-gravatar";

interface GravatarAvatarProps {
  email: string;
  size?: number;
  className?: string;
  alt?: string;
  url?: string; // 新增：自定义头像url
}

const GravatarAvatar: React.FC<GravatarAvatarProps> = ({
  email,
  size = 40,
  className = "",
  alt = "avatar",
  url,
}) => {
  if (url) {
    return (
      <Image
        src={url}
        width={size}
        height={size}
        className={`rounded-full object-cover ${className}`}
        alt={alt}
        referrerPolicy="no-referrer"
      />
    );
  }
  return (
    <Gravatar
      email={email}
      size={size}
      className={`rounded-full ${className}`}
      alt={alt}
      default="identicon"
    />
  );
};

export default GravatarAvatar;