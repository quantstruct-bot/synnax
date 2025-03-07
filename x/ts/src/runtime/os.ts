// Copyright 2024 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { z } from "zod";

export const OPERATING_SYSTEMS = ["MacOS", "Windows", "Linux", "Docker"] as const;
const LOWERCASE_OPERATING_SYSTEMS = ["macos", "windows", "linux", "docker"] as const;
const LOWER_TO_UPPER_OPERATING_SYSTEMS: Record<
  (typeof LOWERCASE_OPERATING_SYSTEMS)[number],
  (typeof OPERATING_SYSTEMS)[number]
> = {
  macos: "MacOS",
  windows: "Windows",
  linux: "Linux",
  docker: "Docker",
};
export const osZ = z
  .enum(OPERATING_SYSTEMS)
  .or(
    z
      .enum(LOWERCASE_OPERATING_SYSTEMS)
      .transform((s) => LOWER_TO_UPPER_OPERATING_SYSTEMS[s]),
  );
export type OS = (typeof OPERATING_SYSTEMS)[number];

export type RequiredGetOSProps = {
  force?: OS;
  default?: OS;
};

export type OptionalGetOSProps = {
  force?: OS | undefined;
  default: OS;
};

export type GetOSProps = RequiredGetOSProps | OptionalGetOSProps;

let os: OS | undefined;

const evalOS = (): OS | undefined => {
  if (typeof window === "undefined") return undefined;
  const userAgent = window.navigator.userAgent.toLowerCase();
  if (userAgent.includes("mac")) return "MacOS";
  if (userAgent.includes("win")) return "Windows";
  if (userAgent.includes("linux")) return "Linux";
  return undefined;
};

export interface GetOS {
  (props?: RequiredGetOSProps): OS;
  (props?: OptionalGetOSProps): OS | undefined;
}

export const getOS = ((props: GetOSProps = {}): OS | undefined => {
  const { force, default: default_ } = props;
  if (force != null) return force;
  if (os != null) return os;
  os = evalOS();
  return os ?? default_;
}) as unknown as GetOS;
