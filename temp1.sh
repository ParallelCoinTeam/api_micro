#!/bin/bash

eval $(minikube docker-env)

KONG_ADMIN="$(minikube service kong-admin --url)"
echo $KONG_ADMIN


#curl -i -X DELETE --url $KONG_ADMIN/consumers/
		curl -i -X DELETE --url $KONG_ADMIN/consumers/33ebb23d-f9a4-4301-b78c-bf1e622595e3
		curl -i -X DELETE --url $KONG_ADMIN/consumers/eb9e6a31-1592-43d8-a9ba-d6dfd8fe815b
		curl -i -X DELETE --url $KONG_ADMIN/consumers/06c017ba-06d8-4d6f-8a9c-73128ab7fc37
		curl -i -X DELETE --url $KONG_ADMIN/consumers/1404c9af-1951-46ea-84e4-04638e460c0c
		curl -i -X DELETE --url $KONG_ADMIN/consumers/373511fd-6415-44b1-813c-1ca2e409f76c
		curl -i -X DELETE --url $KONG_ADMIN/consumers/7c944ec0-3686-4f6e-8f1a-8372dc162a4f
		curl -i -X DELETE --url $KONG_ADMIN/consumers/374a9680-bf8d-491c-8e79-b445c39a0ce9
		curl -i -X DELETE --url $KONG_ADMIN/consumers/20fece67-685f-4c10-a2fe-aeb6ffdf9aac
		curl -i -X DELETE --url $KONG_ADMIN/consumers/0073e16d-2596-425a-84c1-a889334ca41c
		curl -i -X DELETE --url $KONG_ADMIN/consumers/339cbf8c-b86a-4136-9c8f-11f90e3e1408
		curl -i -X DELETE --url $KONG_ADMIN/consumers/45735101-b627-447f-bbbb-02e258b250ff
		curl -i -X DELETE --url $KONG_ADMIN/consumers/78e95a46-d60b-47ba-a9c4-97c239c279a8
		curl -i -X DELETE --url $KONG_ADMIN/consumers/edc50501-40b5-4b1c-8e79-c7689efe1a6d
		curl -i -X DELETE --url $KONG_ADMIN/consumers/1a46e0af-ed0e-4cb7-bddb-fcc8c7faa492
		curl -i -X DELETE --url $KONG_ADMIN/consumers/b2b098bd-b0f6-435d-a371-b50b1710eccb
		curl -i -X DELETE --url $KONG_ADMIN/consumers/4d0cf37a-58e3-4b7a-948d-ad8319e2e1b5
		curl -i -X DELETE --url $KONG_ADMIN/consumers/57f1ef56-3f4f-49b8-8461-e7fdf01c986c
		curl -i -X DELETE --url $KONG_ADMIN/consumers/bc517ae8-220c-4317-83c5-63d61f3f92fd
		curl -i -X DELETE --url $KONG_ADMIN/consumers/dedbbcbf-ef26-4bcf-b6c0-9c1d00a62b0c
		curl -i -X DELETE --url $KONG_ADMIN/consumers/8f4b93e0-e88c-4852-95e1-b45097072d4f
		curl -i -X DELETE --url $KONG_ADMIN/consumers/69e6cb5e-70cc-446e-83b1-67ddda94abfd
		curl -i -X DELETE --url $KONG_ADMIN/consumers/f1eb810a-d4a5-4fb6-8e6c-421238a01f53
		curl -i -X DELETE --url $KONG_ADMIN/consumers/975d3141-d765-47ea-9d83-348e0e2821a4
		curl -i -X DELETE --url $KONG_ADMIN/consumers/a1020d46-585f-4c9a-9b70-58f4c0fdd7d5
