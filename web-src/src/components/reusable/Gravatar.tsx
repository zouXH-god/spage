"use client";

import React from "react";
import Gravatar from "react-gravatar";

import Image from "next/image";

import { User } from "@/api/user.models";

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

export function getGravatarByUser(user?: User, className: string = ""): React.ReactElement {
  if (!user) {
    return <GravatarAvatar email="" />;
  }
  return (
    <GravatarAvatar
      email={user.email || ""}
      size={40}
      className={className}
      alt={user.displayName || user.name}
      url={user.avatarUrl} // 使用用户的自定义头像URL
    />
  );
}

export default GravatarAvatar;
