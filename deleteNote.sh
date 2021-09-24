
#!/bin/sh
path=$(cd "$(dirname "$0")";pwd)
echo ${path}


# 删除 //注释
find ${path} \( -name "*.go" \) -print | xargs sed -ig 's/^[[:space:]]*\/\/.*//g'

# 删除 //注释
find ${path} \( -name "*.go" \) | xargs sed -ig 's/[[:space:]]\/\/.*//g'

# 删除不跨行 /* */
find ${path} \( -name "*.go" \) | xargs sed -ig 's/\/\*.*\*\///g'

# 删除跨行 /* */在行内
find ${path} \( -name "*.go" \) | xargs sed -ig '/\W\/\*/,/\*\//d'

# 删除跨行 /* */在行首
find ${path} \( -name "*.go" \) | xargs sed -ig '/^[[:space:]]*\/\*/,/\*\//d'

# 删除 #pragma
find ${path} \( -name "*.go" \) | xargs sed -ig '/#pragma/d'

# 删除 #warning
find ${path} \( -name "*.go" \) | xargs sed -ig '/#warning/d'

# 删除备份文件
find ${path} \( -name "*.gog" \) | xargs rm
