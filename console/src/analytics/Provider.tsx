import posthog from "posthog-js";
import * as base from "posthog-js";
import { PostHogProvider as BaseProvider } from "posthog-js/react";
import { type PropsWithChildren } from "react";

const POSTHOG_API_KEY = "phc_UilxNDaUWDRcJ2G2aqzQEVPTxObKadiPVIY1eoMAlvx";

const baseClient = posthog.init(POSTHOG_API_KEY, {
  api_host: "https://us.i.posthog.com",
  person_profiles: "always",
});

export interface ProviderProps extends PropsWithChildren<{}> {}

export const Provider = ({ children }: ProviderProps) => (
  <BaseProvider client={baseClient}>{children}</BaseProvider>
);
