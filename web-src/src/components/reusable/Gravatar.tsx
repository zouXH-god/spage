import React from "react";
import Gravatar from "react-gravatar";

interface GravatarAvatarProps {
  email: string;
  size?: number;
  className?: string;
  alt?: string;
}

const GravatarAvatar: React.FC<GravatarAvatarProps> = ({
  email,
  size = 40,
  className = "",
  alt = "avatar",
}) => (
  <Gravatar
    email={email}
    size={size}
    className={`rounded-full ${className}`}
    alt={alt}
    default="identicon"
  />
);

export default GravatarAvatar;