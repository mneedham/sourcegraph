-- This migration was generated by the command `sg telemetry add`
DELETE FROM event_logs_export_allowlist WHERE event_name IN (SELECT * FROM UNNEST('{CodySignup,VSCodeInstall,VSCodeMarketplace,TryCodyWeb,TryCodyWebOnboardingDisplayed}'::TEXT[]));