# Script to update prod version on conf
version=$(git rev-parse --short HEAD)
sed -i'.bak' -e "s/PROD_VERSION/$version/g" conf/conf_prod.go