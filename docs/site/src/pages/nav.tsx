// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

import { clusterNav } from "@/pages/reference/cluster/_nav";
import { conceptsNav } from "@/pages/reference/concepts/_nav";
import { consoleNav } from "@/pages/reference/console/_nav";
import { pythonClientNav } from "@/pages/reference/python-client/_nav";
import { analystNav } from "@/pages/guides/analyst/nav";
import { sysAdminNav } from "@/pages/guides/sys-admin/nav";
import { operationsNav } from "@/pages/guides/operations/nav";
import { typescriptClientNav } from "@/pages/reference/typescript-client/_nav";

export const componentsPages = [
  {
    name: "Get Started",
    key: "/reference/",
    href: "/reference/",
  },
  conceptsNav,
  clusterNav,
  pythonClientNav,
  typescriptClientNav,
  consoleNav,
];
export const guidesPages = [
  {
    name: "Get Started",
    key: "/guides/",
    href: "/guides/",
  },
  analystNav,
  sysAdminNav,
  {/*operationsNav,*/}
];
