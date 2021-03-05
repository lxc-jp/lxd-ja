# Backing up a LXD server
## What to backup
When planning to backup a LXD server, consider all the different objects
that are stored/managed by LXD:

 - Instances (database records and filesystems)
 - Images (database records, image files and filesystems)
 - Networks (database records and state files)
 - Profiles (database records)
 - Storage volumes (database records and filesystems)

Only backing up the database or only backing up the instances will not
get you a fully functional backup.

In some disaster recovery scenarios, that may be reasonable but if your
goal is to get back online quickly, consider all the different pieces of
LXD you're using.

## Full backup
A full backup would include the entirety of `/var/lib/lxd` or
`/var/snap/lxd/common/lxd` for snap users.

You will also need to appropriately backup any external storage that you
made LXD use, this can be LVM volume groups, ZFS zpools or any other
resource which isn't directly self-contained to LXD.

Restoring involves stopping LXD on the target server, wiping the lxd
directory, restoring the backup and any external dependency it requires.

If not using the snap package and your source system has a /etc/subuid
and /etc/subgid file, restoring those or at least the entries inside
them for both the `lxd` and `root` user is also a good idea
(avoids needless shifting of container filesystems).

Then start LXD again and check that everything works fine.

## Secondary backup LXD server
LXD supports copying and moving instances and storage volumes between two hosts.

So with a spare server, you can copy your instances and storage volumes
to that secondary server every so often, allowing it to act as either an
offline spare or just as a storage server that you can copy your
instances back from if needed.

## Instance backups
The `lxc export` command can be used to export instances to a backup tarball.
Those tarballs will include all snapshots by default and an "optimized"
tarball can be obtained if you know that you'll be restoring on a LXD
server using the same storage pool backend.

You can use any compressor installed on the server using the `--compression` 
flag. There is no validation on the LXD side, any command that is available
to LXD and supports `-c` for stdout should work.

Those tarballs can be saved any way you want on any filesystem you want
and can be imported back into LXD using the `lxc import` command.

## Disaster recovery
Additionally, LXD maintains a `backup.yaml` file in each instance's storage
volume. This file contains all necessary information to recover a given
instance, such as instance configuration, attached devices and storage.

This file can be processed by the `lxd import` command, not to
be confused with `lxc import`.

To use the disaster recovery mechanism, you must mount the instance's
storage to its expected location, usually under
`storage-pools/NAME-OF-POOL/containers/NAME-OF-CONTAINER`.

Depending on your storage backend you will also need to do the same for
any snapshot you want to restore (needed for `dir` and `btrfs`).

Once everything is mounted where it should be, you can now run `lxd import NAME-OF-CONTAINER`.

If any matching database entry for resources declared in `backup.yaml` is found
during import, the command will refuse to restore the instance.  This can be
overridden by passing `--force`.

NOTE: When dealing with mounts and the snap, you may need to either
perform a full restart of the snap with `snap stop` and `snap start` or
perform the mounts from within the snap environment using `nsenter
--mount=/run/snapd/ns/lxd.mnt`.
