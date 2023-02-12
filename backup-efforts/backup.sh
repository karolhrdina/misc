#!/bin/bash

# Poor man's pendrive backup tool.
#
# Will back up files and directories specified in an input file to an usb
# mounted pendrive, selecting the proper rsync options depending on the
# filesystem.
# 
# Creates the following structure on target:
#   backup
#   ├── notebook
#
# At the end, copies itself and the input file as well.
#


# get_device_fs(device_name):
#   Returns filesystem of the device, where:
#       'device_name` - full (canonical) path to device without trailing slash
#                       e.g. /run/media/<username>/usb-stick
function get_device_fs() {
    local device_name=$(mount | grep "$1" | awk '{print $1}')
    local fstype_name=$(lsblk -o FSTYPE "${device_name}" | tail -n 1)
    echo "$fstype_name"
}

# usage():
#   Prints script usage.
function usage() {
    echo "Usage: ${0} <target> <input>"
}

# Input check
# TARGET is where we want to create the backup
TARGET="${1}"

if [[ -z "${TARGET}" ]]; then
    echo "ERROR:"
    echo "Please specify a mounted device in /run/media/$(whoami). At the moment, options are:"
    echo "$(ls /run/media/`whoami`)"
    echo
    usage
    exit 1
fi

# INPUT_FILE is a file with newline separated list of what to backup
INPUT_FILE="${2}"

# Get the filesystem of the target.
TARGET=${TARGET%/} # Remove trailing slash, if present
FSTYPE=$(get_device_fs "${TARGET}")
RSYNC_OPTS=""

# If it's `exFAT`, we have to use different rsync options.
#   https://www.scivision.dev/rsync-to-exfat-drive/
#   https://stackoverflow.com/questions/32682694/rsync-backup-to-external-hard-disk-exfat-fails
#
#   rsync options explained:
#       --delete            delete extraneous files from dest dirs
#       --ignore-errors     delete even if there are I/O errors
if [[ "${FSTYPE}" == "exfat" ]]; then
    RSYNC_OPTS="-rltD --delete --ignore-errors"
else
    RSYNC_OPTS="-aX --delete --ignore-errors"
fi

# SUMMARY
echo
echo -e "TARGET:\t\t${TARGET}"
echo -e "FSTYPE:\t\t${FSTYPE}"
echo -e "RSYNC OPTS:\t${RSYNC_OPTS}"

# Ask for user confirmation
read -p "Do you want to continue? (y/N) " yn

case $yn in
	y) ;;
	*) echo "Exiting..."; exit 1;;
esac

# Transform input file lines into command arguments,
# expanding ~ into #$HOME and
# ignoring empty lines, whitespace-only lines
declare -a sources=()

while read -r line && [[ ! -z "${line// }" ]]; do
    line="${line/#\~/$HOME}"
    sources+=(${line})
done < "${INPUT_FILE}"

# Set up the operation itself
RSYNC_TARGET="${TARGET}/backup"
mkdir -p ${RSYNC_TARGET}

set -x
rsync ${RSYNC_OPTS} --stats ${sources[*]} "${RSYNC_TARGET}/notebook"
rsync ${RSYNC_OPTS} "$0" "${INPUT_FILE}" "${RSYNC_TARGET}"
set +x

exit 0
