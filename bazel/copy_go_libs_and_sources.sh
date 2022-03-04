#!/bin/bash

# For go files
base_dir=${1:-$PWD}
# check for empty directory
echo "So we identified the code directory as $base_dir. \
If this is not true then pass the directory to shell script (first command line argument)"
cd $base_dir || exit

IFS=$'\n' read -r -d '' -a go_files < <(find -L ./bazel-bin -iname "*.pb.go" && printf '\0')

for go_file in "${go_files[@]}"
do
#  echo "Go file: $go_file"
  new_file=$go_file
  path=''
#  echo "Dirname: $(dirname $new_file)"
#  echo "Basename: $(basename $new_file)"

  # Find the import path of go code
  while ! [[ "$new_file" == *% || "$new_file" == *_ ]]
  do
      file_part=$(basename $new_file)
      if [[ -z "$path" ]]
      then
        path=$file_part
      else
        path="${file_part}/${path}"
      fi
      new_file=$(dirname $new_file)
#      echo "New file: $new_file"
#      echo "File part: $file_part"
  done

#  echo "Path: $path"
#  echo "Go go_file: $go_file, New file: $new_file"
  # Find base module of this go code
  full_module_name=""
  while ! [[ "$new_file" == *bazel-bin ]]
  do
      module_name=$(basename $new_file)
      full_module_name="$module_name/$full_module_name"
      new_file=$(dirname $new_file)
  done

  while !  [[ -d "./$full_module_name" ]]
  do
    full_module_name=$(dirname $full_module_name)
  done

  if ! [[ "$full_module_name" == . ]]
  then
#    echo "Go Final full module name $full_module_name"
    path="./$full_module_name"
    module_suffix=$(basename $full_module_name)
    required_dir="./$full_module_name/Gen/Go$module_suffix"
    mkdir -p "$required_dir"
    cp -f -R "$go_file" "$required_dir/"
  else
    echo "Ignoring the module: $full_module_name"
  fi
done

