package main

import (
	"fmt"
	"strings"
)

const script = `#!/usr/bin/env bash

runner_script_trap() {
	exit_code=$?
	out_json="{\"command_exit_code\": $exit_code, \"script\": \"$0\"}"

	echo ""
	echo "$out_json"
	exit 0
}

trap runner_script_trap EXIT

set -eo pipefail
set +o noclobber
: | eval $'export FF_CMD_DISABLE_DELAYED_ERROR_LEVEL_EXPANSION=$\'false\'\nexport FF_NETWORK_PER_BUILD=$\'false\'\nexport FF_USE_LEGACY_KUBERNETES_EXECUTION_STRATEGY=$\'false\'\nexport FF_USE_DIRECT_DOWNLOAD=$\'true\'\nexport FF_SKIP_NOOP_BUILD_STAGES=$\'true\'\nexport FF_USE_FASTZIP=$\'false\'\nexport FF_GITLAB_REGISTRY_HELPER_IMAGE=$\'true\'\nexport FF_DISABLE_UMASK_FOR_DOCKER_EXECUTOR=$\'false\'\nexport FF_ENABLE_BASH_EXIT_CODE_CHECK=$\'false\'\nexport FF_USE_WINDOWS_LEGACY_PROCESS_STRATEGY=$\'true\'\nexport FF_USE_NEW_BASH_EVAL_STRATEGY=$\'false\'\nexport FF_USE_POWERSHELL_PATH_RESOLVER=$\'false\'\nexport FF_USE_DYNAMIC_TRACE_FORCE_SEND_INTERVAL=$\'false\'\nexport FF_SCRIPT_SECTIONS=$\'false\'\nexport FF_USE_NEW_SHELL_ESCAPE=$\'false\'\nexport FF_ENABLE_JOB_CLEANUP=$\'false\'\nexport CI_JOB_IMAGE=$\'gitlab-registry.ozon.ru/docker/ci:latest\'\nexport CI_RUNNER_SHORT_TOKEN=$\'tR-GRJHV\'\nexport ARTIFACT_COMPRESSION_LEVEL=$\'default\'\nexport ARTIFACT_DOWNLOAD_ATTEMPTS=3\nexport CACHE_COMPRESSION_LEVEL=$\'fastest\'\nexport DAS_TOKEN=$\'eyJhbGciOiJIUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICI5MGQwY2U1Ni00MTRhLTRlNzQtYWE5MC1jYjc1M2ViMDQ4N2IifQ.eyJpYXQiOjE2Mzg4NjA5MDIsImp0aSI6IjRlZGUxMzEyLTE2YWYtNDBiMi1hZjg5LTIzY2VlZWYzOGMxYyIsImlzcyI6Imh0dHBzOi8vc3NvLm8zLnJ1L2F1dGgvcmVhbG1zL3JlIiwiYXVkIjoiaHR0cHM6Ly9zc28ubzMucnUvYXV0aC9yZWFsbXMvcmUiLCJzdWIiOiJjNzcwMzA3MS1hMzY4LTRiYmUtYTA2Ny1kMTA5ODZmNjFlZWMiLCJ0eXAiOiJPZmZsaW5lIiwiYXpwIjoiZGFzLWFwaSIsInNlc3Npb25fc3RhdGUiOiI5NGY0NDczYy1iOTY1LTQyY2YtYmY0OS04NjQ4MjE3Y2MyYTEiLCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIG9mZmxpbmVfYWNjZXNzIn0.aDZEg6fllD5jbQZE-hEkrebz_fEMCdxVA1T7IJ2qFss\'\nexport DOCKER_CI_IMAGE=$\'gitlab-registry.ozon.ru/docker/ci:latest\'\nexport DOTNET_USE_PRERELEASE_VERSIONS_CHECK=$\'true\'\nexport ENABLE_NEW_NUGET_CONFIG=$\'true\'\nexport ENDPOINTS_ETCD=$\'etcdcommon1-prod.s.o3.ru:2379,etcdcommon2-prod.s.o3.ru:2379,etcdcommon3-prod.s.o3.ru:2379,etcdcommon4-prod.s.o3.ru:2379,etcdcommon5-prod.s.o3.ru:2379\'\nexport FF_ENABLE_BASH_EXIT_CODE_CHECK=$\'true\'\nexport FF_GITLAB_REGISTRY_HELPER_IMAGE=$\'false\'\nexport FF_USE_FASTZIP=$\'true\'\nexport FF_USE_LEGACY_KUBERNETES_EXECUTION_STRATEGY=$\'false\'\nexport GET_SOURCES_ATTEMPTS=3\nexport GITLAB_API_TOKEN=$\'LSLsFcduNcvhHtqMPZEJ\'\nexport GOMAXPROCS=1\nexport GPG_TOKEN=$\'houtoh0ohGoophi\'\nexport K8S_EXECUTOR=$\'true\'\nexport NEXUS_TOKEN=$\'ViraiChah5saich\'\nexport NODE_OPTIONS=$\'--max-old-space-size=4096\'\nexport O3RE_KC_PASS=$\'5iO1MsVRMl4GZK3c\'\nexport SKIP_DELETE_MERGED_TARGET_BRANCHES=$\'true\'\nexport SONAR_HOST_URL=$\'https://sonar.s.o3.ru\'\nexport SONAR_LOGIN=$\'b8e2274af2a342cf90f04e71b3c2dd751e3ac525\'\nexport TRANSFER_METER_FREQUENCY=$\'3s\'\nexport VAULT_TOKEN=$\'s.fHPwQ4YTyhcaViGUF4e0UD9O\'\nexport VT_=$\'5e7qHm5lcNYY8klTdrHF7TBf\'\nexport CI_BUILDS_DIR=$\'/builds\'\nexport CI_PROJECT_DIR=$\'/builds/tR-GRJHV/0/artemikhaylov-test-group/artemikhaylov-dummy-project\'\nexport CI_CONCURRENT_ID=7\nexport CI_CONCURRENT_PROJECT_ID=0\nexport CI_SERVER=$\'yes\'\nexport CI_JOB_STATUS=$\'running\'\nmkdir -p "/builds/tR-GRJHV/0/artemikhaylov-test-group/artemikhaylov-dummy-project.tmp"\necho -n $\'-----BEGIN CERTIFICATE-----\\nMIIGXzCCBUegAwIBAgIMfg20Uv7v3r+RKs07MA0GCSqGSIb3DQEBCwUAMFAxCzAJ\\nBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9iYWxTaWduIG52LXNhMSYwJAYDVQQDEx1H\\nbG9iYWxTaWduIFJTQSBPViBTU0wgQ0EgMjAxODAeFw0yMjA5MTYwOTU2MDZaFw0y\\nMzEwMTgwOTU2MDVaMGQxCzAJBgNVBAYTAlJVMQ8wDQYDVQQIEwZNb3Njb3cxDzAN\\nBgNVBAcTBk1vc2NvdzEfMB0GA1UEChMWSW50ZXJuZXQgU29sdXRpb25zIExMQzES\\nMBAGA1UEAwwJKi5vem9uLnJ1MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC\\nAQEAroprdEKy0PfStG2El8X2sRr+5xjIGUu2LMrVaG98ff7Ouw+0lJ8HlHdlogOW\\nPtpTNH0IQlljSwr5wkxrKWoIhdwVo+f7lpuJ9rDrtD0oEm6OFwWi04JYWtZ7ThNK\\nz0TJS3+Y+LGKK08+EVjkqgpEAg5Zaxw6dilFDal42xyInDUyrAqFhBoq7VNz0yVW\\n2MOm25j2hDak6w2XSsayLk5gFU9wrSenY9C6Y6hd3y8cO2pdfHiiRTum+nmw+38q\\nUsadBfJVtOthjOCCiN9p/JHBPhxCmVDGNilmXFjZ+IIGPRwoIjrnGedM7eqTHD1H\\n/tyNIJ5Q4hXh7JD5TOqaEufJJwIDAQABo4IDIzCCAx8wDgYDVR0PAQH/BAQDAgWg\\nMIGOBggrBgEFBQcBAQSBgTB/MEQGCCsGAQUFBzAChjhodHRwOi8vc2VjdXJlLmds\\nb2JhbHNpZ24uY29tL2NhY2VydC9nc3JzYW92c3NsY2EyMDE4LmNydDA3BggrBgEF\\nBQcwAYYraHR0cDovL29jc3AuZ2xvYmFsc2lnbi5jb20vZ3Nyc2FvdnNzbGNhMjAx\\nODBWBgNVHSAETzBNMEEGCSsGAQQBoDIBFDA0MDIGCCsGAQUFBwIBFiZodHRwczov\\nL3d3dy5nbG9iYWxzaWduLmNvbS9yZXBvc2l0b3J5LzAIBgZngQwBAgIwCQYDVR0T\\nBAIwADA2BgNVHREELzAtggkqLm96b24ucnWCCioub3pvbmUucnWCC3d3dy5vem9u\\nLnJ1ggdvem9uLnJ1MB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNV\\nHSMEGDAWgBT473/yzXhnqN5vjySNiPGHAwKz6zAdBgNVHQ4EFgQUmwVI3pfPH61m\\nrTT284sBAP+q3yQwggGABgorBgEEAdZ5AgQCBIIBcASCAWwBagB3AOg+0No+9QY1\\nMudXKLyJa8kD08vREWvs62nhd31tBr1uAAABg0W6GpgAAAQDAEgwRgIhAIu/LaGC\\neXc4XHVoz4lSzsniA0TJnnkvFotrJM9MzPotAiEAwwLQyGhsd0a8izDRuQ0FU2dR\\nPIrCkeE63E+1hju/TnYAdgBvU3asMfAxGdiZAKRRFf93FRwR2QLBACkGjbIImjfZ\\nEwAAAYNFuhqWAAAEAwBHMEUCIQD2z7ONI91FHmEJEu75hZoQ6xTF1oMuanjzcOBg\\nrFwwgAIgL5dRWyEaMXjcCFgK0KXuXzDPdKj5FgiQyyq6KaH6400AdwCzc3cH4YRQ\\n+GOG1gWp3BEJSnktsWcMC4fc8AMOeTalmgAAAYNFuhq6AAAEAwBIMEYCIQDQMiuf\\nVujlqaDsmz3S8y7JNu9t59cdz1uoT7+Kvby9EQIhALXdeUOOrQqy0mcey0o+ZjaS\\nevSNy3gQsKdzeDrf7ppmMA0GCSqGSIb3DQEBCwUAA4IBAQBUFnhK+ROMOE8g2dUP\\nEtWDRWL45lcDWk30GEuCmharGA8QdqmBtp3ekBJTCmBCwbkp2KobcwJIvFIJb9a4\\ncpmkPX/2wnXMvte94Rjotkif4cQdJQ6oyqxl8BSj8aJhiF6w0c1DfVmbmSAFV+Ml\\nYmJo6/j8kJsCtF99YxXfBO32eyhkBQ2KLsFH5slUGFUleNQbdYNQjca8WAiZ/lfh\\nK5uElRJsavLxei/5GqTkKiqFgbdGNq5f/KZ3E1dLtGp9PX1NQcyLLDSR/Fsx0JcN\\nd7mJfR1AazPYddYhLeL1BOb7R/1wA7jcfCGR/on9jW2oiTYaqIfTxQmV1ZCuDwHL\\nSi/p\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\nMIIETjCCAzagAwIBAgINAe5fIh38YjvUMzqFVzANBgkqhkiG9w0BAQsFADBMMSAw\\nHgYDVQQLExdHbG9iYWxTaWduIFJvb3QgQ0EgLSBSMzETMBEGA1UEChMKR2xvYmFs\\nU2lnbjETMBEGA1UEAxMKR2xvYmFsU2lnbjAeFw0xODExMjEwMDAwMDBaFw0yODEx\\nMjEwMDAwMDBaMFAxCzAJBgNVBAYTAkJFMRkwFwYDVQQKExBHbG9iYWxTaWduIG52\\nLXNhMSYwJAYDVQQDEx1HbG9iYWxTaWduIFJTQSBPViBTU0wgQ0EgMjAxODCCASIw\\nDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKdaydUMGCEAI9WXD+uu3Vxoa2uP\\nUGATeoHLl+6OimGUSyZ59gSnKvuk2la77qCk8HuKf1UfR5NhDW5xUTolJAgvjOH3\\nidaSz6+zpz8w7bXfIa7+9UQX/dhj2S/TgVprX9NHsKzyqzskeU8fxy7quRU6fBhM\\nabO1IFkJXinDY+YuRluqlJBJDrnw9UqhCS98NE3QvADFBlV5Bs6i0BDxSEPouVq1\\nlVW9MdIbPYa+oewNEtssmSStR8JvA+Z6cLVwzM0nLKWMjsIYPJLJLnNvBhBWk0Cq\\no8VS++XFBdZpaFwGue5RieGKDkFNm5KQConpFmvv73W+eka440eKHRwup08CAwEA\\nAaOCASkwggElMA4GA1UdDwEB/wQEAwIBhjASBgNVHRMBAf8ECDAGAQH/AgEAMB0G\\nA1UdDgQWBBT473/yzXhnqN5vjySNiPGHAwKz6zAfBgNVHSMEGDAWgBSP8Et/qC5F\\nJK5NUPpjmove4t0bvDA+BggrBgEFBQcBAQQyMDAwLgYIKwYBBQUHMAGGImh0dHA6\\nLy9vY3NwMi5nbG9iYWxzaWduLmNvbS9yb290cjMwNgYDVR0fBC8wLTAroCmgJ4Yl\\naHR0cDovL2NybC5nbG9iYWxzaWduLmNvbS9yb290LXIzLmNybDBHBgNVHSAEQDA+\\nMDwGBFUdIAAwNDAyBggrBgEFBQcCARYmaHR0cHM6Ly93d3cuZ2xvYmFsc2lnbi5j\\nb20vcmVwb3NpdG9yeS8wDQYJKoZIhvcNAQELBQADggEBAJmQyC1fQorUC2bbmANz\\nEdSIhlIoU4r7rd/9c446ZwTbw1MUcBQJfMPg+NccmBqixD7b6QDjynCy8SIwIVbb\\n0615XoFYC20UgDX1b10d65pHBf9ZjQCxQNqQmJYaumxtf4z1s4DfjGRzNpZ5eWl0\\n6r/4ngGPoJVpjemEuunl1Ig423g7mNA2eymw0lIYkN5SQwCuaifIFJ6GlazhgDEw\\nfpolu4usBCOmmQDo8dIm7A9+O4orkjgTHY+GzYZSR+Y0fFukAj6KYXwidlNalFMz\\nhriSqHKvoflShx8xpfywgVcvzfTO3PYkz6fiNJBonf6q8amaEsybwMbDqKWwIX7e\\nSPY=\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\nMIIDXzCCAkegAwIBAgILBAAAAAABIVhTCKIwDQYJKoZIhvcNAQELBQAwTDEgMB4G\\nA1UECxMXR2xvYmFsU2lnbiBSb290IENBIC0gUjMxEzARBgNVBAoTCkdsb2JhbFNp\\nZ24xEzARBgNVBAMTCkdsb2JhbFNpZ24wHhcNMDkwMzE4MTAwMDAwWhcNMjkwMzE4\\nMTAwMDAwWjBMMSAwHgYDVQQLExdHbG9iYWxTaWduIFJvb3QgQ0EgLSBSMzETMBEG\\nA1UEChMKR2xvYmFsU2lnbjETMBEGA1UEAxMKR2xvYmFsU2lnbjCCASIwDQYJKoZI\\nhvcNAQEBBQADggEPADCCAQoCggEBAMwldpB5BngiFvXAg7aEyiie/QV2EcWtiHL8\\nRgJDx7KKnQRfJMsuS+FggkbhUqsMgUdwbN1k0ev1LKMPgj0MK66X17YUhhB5uzsT\\ngHeMCOFJ0mpiLx9e+pZo34knlTifBtc+ycsmWQ1z3rDI6SYOgxXG71uL0gRgykmm\\nKPZpO/bLyCiR5Z2KYVc3rHQU3HTgOu5yLy6c+9C7v/U9AOEGM+iCK65TpjoWc4zd\\nQQ4gOsC0p6Hpsk+QLjJg6VfLuQSSaGjlOCZgdbKfd/+RFO+uIEn8rUAVSNECMWEZ\\nXriX7613t2Saer9fwRPvm2L7DWzgVGkWqQPabumDk3F2xmmFghcCAwEAAaNCMEAw\\nDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wHQYDVR0OBBYEFI/wS3+o\\nLkUkrk1Q+mOai97i3Ru8MA0GCSqGSIb3DQEBCwUAA4IBAQBLQNvAUKr+yAzv95ZU\\nRUm7lgAJQayzE4aGKAczymvmdLm6AC2upArT9fHxD4q/c2dKg8dEe3jgr25sbwMp\\njjM5RcOO5LlXbKr8EpbsU8Yt5CRsuZRj+9xTaGdWPoO4zzUhw8lo/s7awlOqzJCK\\n6fBdRoyV3XpYKBovHd7NADdBj+1EbddTKJd+82cEHhXXipa0095MJ6RMG3NzdvQX\\nmcIfeg7jLQitChws/zyrVQ4PkX4268NXSb7hLi18YIvDQVETI53O9zJrlAGomecs\\nMx86OyXShkDOOyyGeMlhLxS67ttVb9+E7gUJTb0o2HLO02JQZR7rkpeDMdmztcpH\\nWD9f\\n-----END CERTIFICATE-----\' > "/builds/tR-GRJHV/0/artemikhaylov-test-group/artemikhaylov-dummy-project.tmp/CI_SERVER_TLS_CA_FILE"\nexport CI_SERVER_TLS_CA_FILE="/builds/tR-GRJHV/0/artemikhaylov-test-group/artemikhaylov-dummy-project.tmp/CI_SERVER_TLS_CA_FILE"\nexport CI_PIPELINE_ID=13665970\nexport CI_PIPELINE_URL=$\'https://gitlab.ozon.ru/artemikhaylov-test-group/artemikhaylov-dummy-project/-/pipelines/13665970\'\nexport CI_JOB_ID=121408015\nexport CI_JOB_URL=$\'https://gitlab.ozon.ru/artemikhaylov-test-group/artemikhaylov-dummy-project/-/jobs/121408015\'\nexport CI_JOB_TOKEN=$\'nwys_dJSeN2xyn47VRdD\'\nexport CI_JOB_STARTED_AT=$\'2023-03-29T12:57:52Z\'\nexport CI_BUILD_ID=121408015\nexport CI_BUILD_TOKEN=$\'nwys_dJSeN2xyn47VRdD\'\nexport CI_REGISTRY_USER=$\'gitlab-ci-token\'\nexport CI_REGISTRY_PASSWORD=$\'nwys_dJSeN2xyn47VRdD\'\nexport CI_REPOSITORY_URL=$\'https://gitlab-ci-token:nwys_dJSeN2xyn47VRdD@gitlab.ozon.ru/artemikhaylov-test-group/artemikhaylov-dummy-project.git\'\nexport CI_JOB_JWT=$\'eyJhbGciOiJSUzI1NiIsImtpZCI6InQ4WUgzSE1adTFxNEFHYktUNDY0UlRtU2RWRHpVMlREM19zVElhQ3NONGciLCJ0eXAiOiJKV1QifQ.eyJuYW1lc3BhY2VfaWQiOiI2OTg5IiwibmFtZXNwYWNlX3BhdGgiOiJhcnRlbWlraGF5bG92LXRlc3QtZ3JvdXAiLCJwcm9qZWN0X2lkIjoiMTU5MzgiLCJwcm9qZWN0X3BhdGgiOiJhcnRlbWlraGF5bG92LXRlc3QtZ3JvdXAvYXJ0ZW1pa2hheWxvdi1kdW1teS1wcm9qZWN0IiwidXNlcl9pZCI6IjQ4NTMiLCJ1c2VyX2xvZ2luIjoiYXJ0ZW1pa2hheWxvdiIsInVzZXJfZW1haWwiOiJhcnRlbWlraGF5bG92QG96b24ucnUiLCJwaXBlbGluZV9pZCI6IjEzNjY1OTcwIiwicGlwZWxpbmVfc291cmNlIjoicHVzaCIsImpvYl9pZCI6IjEyMTQwODAxNSIsInJlZiI6InJlbGVhc2UvUkVURVNULTE1NTciLCJyZWZfdHlwZSI6ImJyYW5jaCIsInJlZl9wcm90ZWN0ZWQiOiJmYWxzZSIsImp0aSI6Ijg1NTQ4ZjEzLTcwMTktNGZkYy1hMmU1LWUzNDRiMDkwNTFhZCIsImlzcyI6ImdpdGxhYi5vem9uLnJ1IiwiaWF0IjoxNjgwMDk0NjcyLCJuYmYiOjE2ODAwOTQ2NjcsImV4cCI6MTY4MDA5NDk3Miwic3ViIjoiam9iXzEyMTQwODAxNSJ9.KXkBtwEmEB6Edk_Y32DIYMobEmK13y6wgSEzy6LvtgxeZnjQVFAltp2L-PejOiDQNWJmRJ98gOoYl0NxCRfuwFfCn8iIMUZnpXGMbXQhzqDUxQsKrXIRw6hDKYE4q2qec32bpz88DlnBoHHdJM_3DcM9FvewrMbeMGb9-st0hqn_wxfPT_OxIwN3AtLr_TrryGnTCE4CYHHSIZf61cDkgSFhqPPZ9GukTB42rbs5MV2BGDl3F_VjC7Nsb6UGw2QpreUUgC6ez1Ietjck3gAO56bhsT4VSipCPkm_muLXOhX1HP2VKU1Z2oHJlaS1Hy2qg1telFhPM7R-WPIPJ_Grng\'\nexport CI_JOB_JWT_V1=$\'eyJhbGciOiJSUzI1NiIsImtpZCI6InQ4WUgzSE1adTFxNEFHYktUNDY0UlRtU2RWRHpVMlREM19zVElhQ3NONGciLCJ0eXAiOiJKV1QifQ.eyJuYW1lc3BhY2VfaWQiOiI2OTg5IiwibmFtZXNwYWNlX3BhdGgiOiJhcnRlbWlraGF5bG92LXRlc3QtZ3JvdXAiLCJwcm9qZWN0X2lkIjoiMTU5MzgiLCJwcm9qZWN0X3BhdGgiOiJhcnRlbWlraGF5bG92LXRlc3QtZ3JvdXAvYXJ0ZW1pa2hheWxvdi1kdW1teS1wcm9qZWN0IiwidXNlcl9pZCI6IjQ4NTMiLCJ1c2VyX2xvZ2luIjoiYXJ0ZW1pa2hheWxvdiIsInVzZXJfZW1haWwiOiJhcnRlbWlraGF5bG92QG96b24ucnUiLCJwaXBlbGluZV9pZCI6IjEzNjY1OTcwIiwicGlwZWxpbmVfc291cmNlIjoicHVzaCIsImpvYl9pZCI6IjEyMTQwODAxNSIsInJlZiI6InJlbGVhc2UvUkVURVNULTE1NTciLCJyZWZfdHlwZSI6ImJyYW5jaCIsInJlZl9wcm90ZWN0ZWQiOiJmYWxzZSIsImp0aSI6Ijg1NTQ4ZjEzLTcwMTktNGZkYy1hMmU1LWUzNDRiMDkwNTFhZCIsImlzcyI6ImdpdGxhYi5vem9uLnJ1IiwiaWF0IjoxNjgwMDk0NjcyLCJuYmYiOjE2ODAwOTQ2NjcsImV4cCI6MTY4MDA5NDk3Miwic3ViIjoiam9iXzEyMTQwODAxNSJ9.KXkBtwEmEB6Edk_Y32DIYMobEmK13y6wgSEzy6LvtgxeZnjQVFAltp2L-PejOiDQNWJmRJ98gOoYl0NxCRfuwFfCn8iIMUZnpXGMbXQhzqDUxQsKrXIRw6hDKYE4q2qec32bpz88DlnBoHHdJM_3DcM9FvewrMbeMGb9-st0hqn_wxfPT_OxIwN3AtLr_TrryGnTCE4CYHHSIZf61cDkgSFhqPPZ9GukTB42rbs5MV2BGDl3F_VjC7Nsb6UGw2QpreUUgC6ez1Ietjck3gAO56bhsT4VSipCPkm_muLXOhX1HP2VKU1Z2oHJlaS1Hy2qg1telFhPM7R-WPIPJ_Grng\'\nexport CI_JOB_JWT_V2=$\'eyJhbGciOiJSUzI1NiIsImtpZCI6InQ4WUgzSE1adTFxNEFHYktUNDY0UlRtU2RWRHpVMlREM19zVElhQ3NONGciLCJ0eXAiOiJKV1QifQ.eyJuYW1lc3BhY2VfaWQiOiI2OTg5IiwibmFtZXNwYWNlX3BhdGgiOiJhcnRlbWlraGF5bG92LXRlc3QtZ3JvdXAiLCJwcm9qZWN0X2lkIjoiMTU5MzgiLCJwcm9qZWN0X3BhdGgiOiJhcnRlbWlraGF5bG92LXRlc3QtZ3JvdXAvYXJ0ZW1pa2hheWxvdi1kdW1teS1wcm9qZWN0IiwidXNlcl9pZCI6IjQ4NTMiLCJ1c2VyX2xvZ2luIjoiYXJ0ZW1pa2hheWxvdiIsInVzZXJfZW1haWwiOiJhcnRlbWlraGF5bG92QG96b24ucnUiLCJwaXBlbGluZV9pZCI6IjEzNjY1OTcwIiwicGlwZWxpbmVfc291cmNlIjoicHVzaCIsImpvYl9pZCI6IjEyMTQwODAxNSIsInJlZiI6InJlbGVhc2UvUkVURVNULTE1NTciLCJyZWZfdHlwZSI6ImJyYW5jaCIsInJlZl9wcm90ZWN0ZWQiOiJmYWxzZSIsImp0aSI6IjNhOGIwZmRjLTI3OWQtNGMyOS1hN2JmLWExNWU5YTU0YjQxOSIsImlzcyI6Imh0dHBzOi8vZ2l0bGFiLm96b24ucnUiLCJpYXQiOjE2ODAwOTQ2NzIsIm5iZiI6MTY4MDA5NDY2NywiZXhwIjoxNjgwMDk0OTcyLCJzdWIiOiJwcm9qZWN0X3BhdGg6YXJ0ZW1pa2hheWxvdi10ZXN0LWdyb3VwL2FydGVtaWtoYXlsb3YtZHVtbXktcHJvamVjdDpyZWZfdHlwZTpicmFuY2g6cmVmOnJlbGVhc2UvUkVURVNULTE1NTciLCJhdWQiOiJodHRwczovL2dpdGxhYi5vem9uLnJ1In0.Kzj-m0LWDp0bgNXU0Wid12l-TVe43WNgKfk7CCkggZZMh3Ty4633htBXrRTKUZ4XEyQu5kSq5JxfXgofpx2-n1E2vzQTCmYwZDlu6YuP58HXR6yyz0ARK1fU-9L3QFbn-k3dgkP1gNZvO50tVnE8V4h0yb-pc3uCc7B6L5X1NKHR-Qk3imRPSsUqoYnn-fHes3r4WRR8FkN_5SJYpdDsYBDRV97smwscrSkgo9gc3C9Q9kV9K6BaSoz2MPbJr6gdMnqnz_vfM-hu53mB7EWqbjVuPDMJ2FlvUx1NQHvYiFq8u65iP_4_pEMnByNMChAoUYAYZM2JjZ8KAh7GVYF4cw\'\nexport CI_JOB_NAME=$\'error\'\nexport CI_JOB_STAGE=$\'.pre\'\nexport CI_NODE_TOTAL=1\nexport CI_BUILD_NAME=$\'error\'\nexport CI_BUILD_STAGE=$\'.pre\'\nexport CI=$\'true\'\nexport GITLAB_CI=$\'true\'\nexport CI_SERVER_URL=$\'https://gitlab.ozon.ru\'\nexport CI_SERVER_HOST=$\'gitlab.ozon.ru\'\nexport CI_SERVER_PORT=443\nexport CI_SERVER_PROTOCOL=$\'https\'\nexport CI_SERVER_NAME=$\'GitLab\'\nexport CI_SERVER_VERSION=$\'15.1.0-ee\'\nexport CI_SERVER_VERSION_MAJOR=15\nexport CI_SERVER_VERSION_MINOR=1\nexport CI_SERVER_VERSION_PATCH=0\nexport CI_SERVER_REVISION=$\'31c24d2d864\'\nexport GITLAB_FEATURES=$\'audit_events,blocked_issues,board_iteration_lists,code_owners,code_review_analytics,contribution_analytics,description_diffs,elastic_search,full_codequality_report,group_activity_analytics,group_bulk_edit,group_webhooks,issuable_default_templates,issue_weights,iterations,ldap_group_sync,member_lock,merge_request_approvers,milestone_charts,multiple_issue_assignees,multiple_ldap_servers,multiple_merge_request_assignees,multiple_merge_request_reviewers,project_merge_request_analytics,protected_refs_for_users,push_rules,repository_mirrors,resource_access_token,seat_link,scoped_issue_board,usage_quotas,visual_review_app,wip_limits,send_emails_from_admin_area,repository_size_limit,adjourned_deletion_for_projects_and_groups,admin_audit_log,auditor_user,blocking_merge_requests,board_assignee_lists,board_milestone_lists,ci_cd_projects,ci_secrets_management,cluster_agents_gitops,cluster_agents_ci_impersonation,cluster_deployments,code_owner_approval_required,commit_committer_check,compliance_framework,custom_compliance_frameworks,cross_project_pipelines,custom_file_templates,custom_file_templates_for_namespace,custom_project_templates,cycle_analytics_for_groups,cycle_analytics_for_projects,db_load_balancing,default_branch_protection_restriction_in_groups,default_project_deletion_protection,disable_name_update_for_users,email_additional_text,epics,extended_audit_events,external_authorization_service_api_management,feature_flags_related_issues,feature_flags_code_references,file_locks,geo,generic_alert_fingerprinting,git_two_factor_enforcement,github_integration,group_allowed_email_domains,group_coverage_reports,group_forking_protection,group_merge_request_analytics,group_milestone_project_releases,group_project_templates,group_repository_analytics,group_saml,group_scoped_ci_variables,group_wikis,incident_sla,incident_metric_upload,ide_schema_config,issues_analytics,jira_issues_integration,ldap_group_sync_filter,merge_pipelines,merge_request_performance_metrics,admin_merge_request_approvers_rules,merge_trains,metrics_reports,multiple_alert_http_integrations,multiple_approval_rules,multiple_group_issue_boards,multiple_iteration_cadences,object_storage,operations_dashboard,package_forwarding,pages_size_limit,password_complexity,productivity_analytics,project_aliases,protected_environments,reject_unsigned_commits,saml_group_sync,scoped_labels,smartcard_auth,swimlanes,type_of_work_analytics,minimal_access_role,unprotection_restrictions,ci_project_subscriptions,incident_timeline_view,oncall_schedules,escalation_policies,export_user_permissions,zentao_issues_integration,coverage_check_approval_rule,issuable_resource_links,group_ip_restriction\'\nexport CI_PROJECT_ID=15938\nexport CI_PROJECT_NAME=$\'artemikhaylov-dummy-project\'\nexport CI_PROJECT_TITLE=$\'artemikhaylov-dummy-project\'\nexport CI_PROJECT_DESCRIPTION=\'\'\nexport CI_PROJECT_PATH=$\'artemikhaylov-test-group/artemikhaylov-dummy-project\'\nexport CI_PROJECT_PATH_SLUG=$\'artemikhaylov-test-group-artemikhaylov-dummy-project\'\nexport CI_PROJECT_NAMESPACE=$\'artemikhaylov-test-group\'\nexport CI_PROJECT_ROOT_NAMESPACE=$\'artemikhaylov-test-group\'\nexport CI_PROJECT_URL=$\'https://gitlab.ozon.ru/artemikhaylov-test-group/artemikhaylov-dummy-project\'\nexport CI_PROJECT_VISIBILITY=$\'internal\'\nexport CI_PROJECT_REPOSITORY_LANGUAGES=$\'makefile,go,dockerfile,shell\'\nexport CI_PROJECT_CLASSIFICATION_LABEL=\'\'\nexport CI_DEFAULT_BRANCH=$\'master\'\nexport CI_CONFIG_PATH=$\'.gitlab-ci.yml\'\nexport CI_PAGES_DOMAIN=$\'gp.o3.ru\'\nexport CI_PAGES_URL=$\'http://artemikhaylov-test-group.gp.o3.ru/artemikhaylov-dummy-project\'\nexport CI_REGISTRY=$\'gitlab-registry.ozon.ru\'\nexport CI_REGISTRY_IMAGE=$\'gitlab-registry.ozon.ru/artemikhaylov-test-group/artemikhaylov-dummy-project\'\nexport CI_API_V4_URL=$\'https://gitlab.ozon.ru/api/v4\'\nexport CI_PIPELINE_IID=1539\nexport CI_PIPELINE_SOURCE=$\'push\'\nexport CI_PIPELINE_CREATED_AT=$\'2023-03-29T12:57:50Z\'\nexport CI_COMMIT_SHA=$\'08ee3c083e8c5d3dbed8caf065acee8c43c1468d\'\nexport CI_COMMIT_SHORT_SHA=$\'08ee3c08\'\nexport CI_COMMIT_BEFORE_SHA=$\'167d6dee4589ba77c89ec843ca7a4f14bb9cf2b5\'\nexport CI_COMMIT_REF_NAME=$\'release/RETEST-1557\'\nexport CI_COMMIT_REF_SLUG=$\'release-retest-1557\'\nexport CI_COMMIT_BRANCH=$\'release/RETEST-1557\'\nexport CI_COMMIT_MESSAGE=$\'[RETEST-1557] test\\n\'\nexport CI_COMMIT_TITLE=$\'[RETEST-1557] test\'\nexport CI_COMMIT_DESCRIPTION=\'\'\nexport CI_COMMIT_REF_PROTECTED=$\'false\'\nexport CI_COMMIT_TIMESTAMP=$\'2023-03-29T15:57:40+03:00\'\nexport CI_COMMIT_AUTHOR=$\'artemikhaylov <artemikhaylov@ozon.ru>\'\nexport CI_BUILD_REF=$\'08ee3c083e8c5d3dbed8caf065acee8c43c1468d\'\nexport CI_BUILD_BEFORE_SHA=$\'167d6dee4589ba77c89ec843ca7a4f14bb9cf2b5\'\nexport CI_BUILD_REF_NAME=$\'release/RETEST-1557\'\nexport CI_BUILD_REF_SLUG=$\'release-retest-1557\'\nexport CI_OPEN_MERGE_REQUESTS=$\'artemikhaylov-test-group/artemikhaylov-dummy-project!103\'\nexport CI_RUNNER_ID=1058\nexport CI_RUNNER_DESCRIPTION=$\'gitlab-ci.k8s.infra-ts - k8s ci-jobs2\'\nexport CI_RUNNER_TAGS=$\'k8s, ci-jobs, true-k8s\'\nexport COMMON_PIPELINE_VERSION=$\'go-0.0.5\'\nexport DOCKER_CI_IMAGE_TAG=$\'latest\'\nexport DOCKER_CI_IMAGE=$\'gitlab-registry.ozon.ru/docker/ci:latest\'\nexport O3_RE_ENABLE_SENTRY=$\'true\'\nexport SCHEMA_MIGRATION=$\'true\'\nexport GIT_DEPTH=3\nexport GIT_FETCH_EXTRA_FLAGS=$\'--no-tags\'\nexport GO_VERSION=1.19\nexport LINT_VERSION=1.49.0\nexport BASE_IMAGE_REPO=$\'gitlab-registry.ozon.ru\'\nexport LINT_IMAGE=$\'gitlab-registry.ozon.ru/platform/lint/1.49.0/platform:v0.10\'\nexport RUNNER_ENV_TAG=$\'release-runner\'\nexport DEPLOY_CONFIG=$\'yes\'\nexport CANARY=$\'yes\'\nexport BLUE_GREEN=$\'yes\'\nexport K8S_NAMESPACE=$\'artemikhaylov-test-group\'\nexport SERVICE_NAME=$\'artemikhaylov-dummy-project\'\nexport DOCKERFILE_PATH=$\'.o3/build/package/Dockerfile\'\nexport GITLAB_USER_ID=4853\nexport GITLAB_USER_EMAIL=$\'artemikhaylov@ozon.ru\'\nexport GITLAB_USER_LOGIN=$\'artemikhaylov\'\nexport GITLAB_USER_NAME=$\'Mikhaylov Artem Yuryevich\'\nexport SPREAD_CONSTRAINTS_PROD=$\'true\'\nexport ENABLE_CODE_DIFF_COVERAGE=$\'true\'\nexport SPREAD_CONSTRAINTS=$\'true\'\nexport GITLAB_DEFAULT_BRANCH=$\'master\'\nexport FASTLANE_OPT_OUT_USAGE=YES\nexport HELM_MULTI_DEPLOYMENT_DATA_CENTERS=$\'z501=1;z502=1;z503=1\'\nexport GIT_MERGE_DEPTH=20\nexport SPREAD_CONSTRAINTS_INFRA=$\'true\'\nexport CHECK_PREBUILD_PROD_SCRIPT=$\'yes\'\nexport CI_FUNC_DEBUG1=$\'true\'\nexport CI_LOG_LEVEL=$\'debug\'\nexport INGRESS_CHECK_TIMEOUT_SEC_STG=90\nexport ITC_ID=8291\nexport O3_RE_ENABLE_SENTRY=$\'false\'\nexport SAVE_RE_VARIABLES=$\'true\'\nexport CI_DISPOSABLE_ENVIRONMENT=$\'true\'\nexport CI_RUNNER_VERSION=14.3.4\nexport CI_RUNNER_REVISION=$\'77516d85\'\nexport CI_RUNNER_EXECUTABLE_ARCH=$\'linux/amd64\'\n$\'cd\' "/builds/tR-GRJHV/0/artemikhaylov-test-group/artemikhaylov-dummy-project"\necho $\'\\x1b[32;1m$ # pre_build_script # collapsed multi-line command\\x1b[0;m\'\n# pre_build_script\nif ! ( curl -sSf --retry 10 --retry-delay 1 --retry-connrefused --header "PRIVATE-TOKEN:${GITLAB_API_TOKEN}" \'https://gitlab.ozon.ru/api/v4/projects/534/repository/files/list-ci-job/raw?ref=0.0.5\' | grep -qE "^${CI_JOB_NAME}$" )\nthen\n  echo "Job ${CI_JOB_NAME} not allowed to run on this runner."\n  echo "Please remove \'tags: [ci-jobs]\' from your pipeline for this job."\n  exit 1\nfi\necho $\'\\x1b[32;1m$ export PIPE_VERSION="0.0.5"\\x1b[0;m\'\nexport PIPE_VERSION="0.0.5"\necho $\'\\x1b[32;1m$ export CI_VERSION="${CI_VERSION:-1.5.x}"\\x1b[0;m\'\nexport CI_VERSION="${CI_VERSION:-1.5.x}"\necho $\'\\x1b[32;1m$ echo "$CI_VERSION"\\x1b[0;m\'\necho "$CI_VERSION"\necho $\'\\x1b[32;1m$ eval "$(curl -m 60 -s https://gitlab.ozon.ru/deploy/ci/-/raw/$CI_VERSION/functions.sh)"\\x1b[0;m\'\neval "$(curl -m 60 -s https://gitlab.ozon.ru/deploy/ci/-/raw/$CI_VERSION/functions.sh)"\necho $\'\\x1b[32;1m$ check_dascli_ver\\x1b[0;m\'\ncheck_dascli_ver\necho $\'\\x1b[32;1m$ sleep 3600\\x1b[0;m\'\nsleep 3600\n'
exit 0`

func main() {
	lines := make([]string, 0)
	for _, line := range strings.Split(script, `\n`) {
		lines = append(lines, line)
	}
	fmt.Println(lines)
}