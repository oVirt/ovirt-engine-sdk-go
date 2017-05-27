# -*- coding: utf-8 -*-

#
# Copyright (c) 2016 Red Hat, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#


from ovirtsdk4 import Error
from ovirtsdk4 import types
from ovirtsdk4.service import Service
from ovirtsdk4.writer import Writer


class AffinityGroupService(Service):
    """
    This service manages a single affinity group.

    """

    def __init__(self, connection, path):
        super(AffinityGroupService, self).__init__(connection, path)
        self._vms_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieve the affinity group details.
        [source,xml]
        ----
        <affinity_group id="00000000-0000-0000-0000-000000000000">
          <name>AF_GROUP_001</name>
          <cluster id="00000000-0000-0000-0000-000000000000"/>
          <positive>true</positive>
          <enforcing>true</enforcing>
        </affinity_group>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove the affinity group.
        [source]
        ----
        DELETE /ovirt-engine/api/clusters/000-000/affinitygroups/123-456
        ----


        This method supports the following parameters:

        `async`:: Indicates if the removal should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        group,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Update the affinity group.


        This method supports the following parameters:

        `group`:: The affinity group.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('group', group, types.AffinityGroup),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(group, headers, query, wait)

    def vms_service(self):
        """
        Returns a reference to the service that manages the
        list of all virtual machines attached to this affinity
        group.

        """
        return AffinityGroupVmsService(self._connection, '%s/vms' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'vms':
            return self.vms_service()
        if path.startswith('vms/'):
            return self.vms_service().service(path[4:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AffinityGroupService:%s' % self._path


class AffinityGroupVmService(Service):
    """
    This service manages a single virtual machine to affinity group assignment.

    """

    def __init__(self, connection, path):
        super(AffinityGroupVmService, self).__init__(connection, path)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove this virtual machine from the affinity group.


        This method supports the following parameters:

        `async`:: Indicates if the removal should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AffinityGroupVmService:%s' % self._path


class AffinityGroupVmsService(Service):
    """
    This service manages a collection of all the virtual machines assigned to an affinity group.

    """

    def __init__(self, connection, path):
        super(AffinityGroupVmsService, self).__init__(connection, path)
        self._vm_service = None

    def add(
        self,
        vm,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add a virtual machine to the affinity group.
        For example, to add the virtual machine 000-000 to affinity group 123-456 send a request to:
        [source]
        ----
        POST /ovirt-engine/api/clusters/000-000/affinitygroups/123-456/vms
        ----
        With the following body:
        [source,xml]
        ----
        <vm id="000-000"/>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('vm', vm, types.Vm),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(vm, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List all virtual machines assigned to this affinity group.


        This method supports the following parameters:

        `max`:: Sets the maximum number of virtual machines to return. If not specified, all the virtual machines are
        returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def vm_service(self, id):
        """
        Access the service that manages the virtual machine assignment to this affinity group.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return AffinityGroupVmService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.vm_service(path)
        return self.vm_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AffinityGroupVmsService:%s' % self._path


class AffinityGroupsService(Service):
    """
    The affinity groups service manages virtual machine relationships and dependencies.

    """

    def __init__(self, connection, path):
        super(AffinityGroupsService, self).__init__(connection, path)
        self._group_service = None

    def add(
        self,
        group,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Create a new affinity group.
        Post a request like in the example below to create a new affinity group:
        [source]
        ----
        POST /ovirt-engine/api/clusters/000-000/affinitygroups
        ----
        And use the following example in its body:
        [source,xml]
        ----
        <affinity_group>
          <name>AF_GROUP_001</name>
          <positive>true</positive>
          <enforcing>true</enforcing>
        </affinity_group>
        ----


        This method supports the following parameters:

        `group`:: The affinity group object to create.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('group', group, types.AffinityGroup),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(group, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List existing affinity groups.


        This method supports the following parameters:

        `max`:: Sets the maximum number of affinity groups to return. If not specified all the affinity groups are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def group_service(self, id):
        """
        Access the affinity group service that manages the affinity group specified by an ID.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return AffinityGroupService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.group_service(path)
        return self.group_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AffinityGroupsService:%s' % self._path


class AffinityLabelService(Service):
    """
    The details of a single affinity label.

    """

    def __init__(self, connection, path):
        super(AffinityLabelService, self).__init__(connection, path)
        self._hosts_service = None
        self._vms_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the details of a label.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a label from the system and clears all assignments
        of the removed label.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        label,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates a label. This call will update all metadata, such as the name
        or description.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('label', label, types.AffinityLabel),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_update(label, headers, query, wait)

    def hosts_service(self):
        """
        List all hosts with this label.

        """
        return AffinityLabelHostsService(self._connection, '%s/hosts' % self._path)

    def vms_service(self):
        """
        List all virtual machines with this label.

        """
        return AffinityLabelVmsService(self._connection, '%s/vms' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'hosts':
            return self.hosts_service()
        if path.startswith('hosts/'):
            return self.hosts_service().service(path[6:])
        if path == 'vms':
            return self.vms_service()
        if path.startswith('vms/'):
            return self.vms_service().service(path[4:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AffinityLabelService:%s' % self._path


class AffinityLabelHostService(Service):
    """
    This service represents a host that has a specific
    label when accessed through the affinitylabels/hosts
    subcollection.

    """

    def __init__(self, connection, path):
        super(AffinityLabelHostService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves details about a host that has this label assigned.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove a label from a host.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AffinityLabelHostService:%s' % self._path


class AffinityLabelHostsService(Service):
    """
    This service represents list of hosts that have a specific
    label when accessed through the affinitylabels/hosts
    subcollection.

    """

    def __init__(self, connection, path):
        super(AffinityLabelHostsService, self).__init__(connection, path)
        self._host_service = None

    def add(
        self,
        host,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add a label to a host.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('host', host, types.Host),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(host, headers, query, wait)

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List all hosts with the label.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def host_service(self, id):
        """
        A link to the specific label-host assignment to
        allow label removal.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return AffinityLabelHostService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.host_service(path)
        return self.host_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AffinityLabelHostsService:%s' % self._path


class AffinityLabelVmService(Service):
    """
    This service represents a vm that has a specific
    label when accessed through the affinitylabels/vms
    subcollection.

    """

    def __init__(self, connection, path):
        super(AffinityLabelVmService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves details about a vm that has this label assigned.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove a label from a vm.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AffinityLabelVmService:%s' % self._path


class AffinityLabelVmsService(Service):
    """
    This service represents list of vms that have a specific
    label when accessed through the affinitylabels/vms
    subcollection.

    """

    def __init__(self, connection, path):
        super(AffinityLabelVmsService, self).__init__(connection, path)
        self._vm_service = None

    def add(
        self,
        vm,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add a label to a vm.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('vm', vm, types.Vm),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(vm, headers, query, wait)

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List all vms with the label.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def vm_service(self, id):
        """
        A link to the specific label-vm assignment to
        allow label removal.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return AffinityLabelVmService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.vm_service(path)
        return self.vm_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AffinityLabelVmsService:%s' % self._path


class AffinityLabelsService(Service):
    """
    Manages the affinity labels available in the system.

    """

    def __init__(self, connection, path):
        super(AffinityLabelsService, self).__init__(connection, path)
        self._label_service = None

    def add(
        self,
        label,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new label. The label is automatically attached
        to all entities mentioned in the vms or hosts lists.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('label', label, types.AffinityLabel),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(label, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists all labels present in the system.


        This method supports the following parameters:

        `max`:: Sets the maximum number of labels to return. If not specified all the labels are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def label_service(self, id):
        """
        Link to a single label details.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return AffinityLabelService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.label_service(path)
        return self.label_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AffinityLabelsService:%s' % self._path


class AreaService(Service):
    """
    This annotation is intended to specify what oVirt area is the annotated concept related to. Currently the following
    areas are in use, and they are closely related to the oVirt teams, but not necessarily the same:
    - Infrastructure
    - Network
    - SLA
    - Storage
    - Virtualization
    A concept may be associated to more than one area, or to no area.
    The value of this annotation is intended for reporting only, and it doesn't affect at all the generated code or the
    validity of the model

    """

    def __init__(self, connection, path):
        super(AreaService, self).__init__(connection, path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AreaService:%s' % self._path


class AssignedAffinityLabelService(Service):
    """
    This service represents one label to entity assignment
    when accessed using the entities/affinitylabels subcollection.

    """

    def __init__(self, connection, path):
        super(AssignedAffinityLabelService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves details about the attached label.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the label from an entity. Does not touch the label itself.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AssignedAffinityLabelService:%s' % self._path


class AssignedAffinityLabelsService(Service):
    """
    This service is used to list and manipulate affinity labels that are
    assigned to supported entities when accessed using entities/affinitylabels.

    """

    def __init__(self, connection, path):
        super(AssignedAffinityLabelsService, self).__init__(connection, path)
        self._label_service = None

    def add(
        self,
        label,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Attaches a label to an entity.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('label', label, types.AffinityLabel),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(label, headers, query, wait)

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists all labels that are attached to an entity.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def label_service(self, id):
        """
        Link to the specific entity-label assignment to allow
        removal.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return AssignedAffinityLabelService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.label_service(path)
        return self.label_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AssignedAffinityLabelsService:%s' % self._path


class AssignedCpuProfileService(Service):
    """
    """

    def __init__(self, connection, path):
        super(AssignedCpuProfileService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AssignedCpuProfileService:%s' % self._path


class AssignedCpuProfilesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(AssignedCpuProfilesService, self).__init__(connection, path)
        self._profile_service = None

    def add(
        self,
        profile,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('profile', profile, types.CpuProfile),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(profile, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def profile_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return AssignedCpuProfileService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.profile_service(path)
        return self.profile_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AssignedCpuProfilesService:%s' % self._path


class AssignedDiskProfileService(Service):
    """
    """

    def __init__(self, connection, path):
        super(AssignedDiskProfileService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AssignedDiskProfileService:%s' % self._path


class AssignedDiskProfilesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(AssignedDiskProfilesService, self).__init__(connection, path)
        self._profile_service = None

    def add(
        self,
        profile,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('profile', profile, types.DiskProfile),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(profile, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def profile_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return AssignedDiskProfileService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.profile_service(path)
        return self.profile_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AssignedDiskProfilesService:%s' % self._path


class AssignedNetworkService(Service):
    """
    """

    def __init__(self, connection, path):
        super(AssignedNetworkService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        network,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('network', network, types.Network),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(network, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AssignedNetworkService:%s' % self._path


class AssignedNetworksService(Service):
    """
    """

    def __init__(self, connection, path):
        super(AssignedNetworksService, self).__init__(connection, path)
        self._network_service = None

    def add(
        self,
        network,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('network', network, types.Network),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(network, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def network_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return AssignedNetworkService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.network_service(path)
        return self.network_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AssignedNetworksService:%s' % self._path


class AssignedPermissionsService(Service):
    """
    Represents a permission sub-collection, scoped by user, group or some entity type.

    """

    def __init__(self, connection, path):
        super(AssignedPermissionsService, self).__init__(connection, path)
        self._permission_service = None

    def add(
        self,
        permission,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Assign a new permission to a user or group for specific entity.
        For example, to assign the `UserVmManager` role to the virtual machine with id `123` to the user with id `456`
        send a request like this:
        ....
        POST /ovirt-engine/api/vms/123/permissions
        ....
        With a request body like this:
        [source,xml]
        ----
        <permission>
          <role>
            <name>UserVmManager</name>
          </role>
          <user id="456"/>
        </permission>
        ----
        To assign the `SuperUser` role to the system to the user with id `456` send a request like this:
        ....
        POST /ovirt-engine/api/permissions
        ....
        With a request body like this:
        [source,xml]
        ----
        <permission>
          <role>
            <name>SuperUser</name>
          </role>
          <user id="456"/>
        </permission>
        ----
        If you want to assign permission to the group instead of the user please replace the `user` element with the
        `group` element with proper `id` of the group. For example to assign the `UserRole` role to the cluster with
        id `123` to the group with id `789` send a request like this:
        ....
        POST /ovirt-engine/api/clusters/123/permissions
        ....
        With a request body like this:
        [source,xml]
        ----
        <permission>
          <role>
            <name>UserRole</name>
          </role>
          <group id="789"/>
        </permission>
        ----


        This method supports the following parameters:

        `permission`:: The permission.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('permission', permission, types.Permission),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(permission, headers, query, wait)

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List all the permissions of the specific entity.
        For example to list all the permissions of the cluster with id `123` send a request like this:
        ....
        GET /ovirt-engine/api/clusters/123/permissions
        ....
        [source,xml]
        ----
        <permissions>
          <permission id="456">
            <cluster id="123"/>
            <role id="789"/>
            <user id="451"/>
          </permission>
          <permission id="654">
            <cluster id="123"/>
            <role id="789"/>
            <group id="127"/>
          </permission>
        </permissions>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def permission_service(self, id):
        """
        Sub-resource locator method, returns individual permission resource on which the remainder of the URI is
        dispatched.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return PermissionService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.permission_service(path)
        return self.permission_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AssignedPermissionsService:%s' % self._path


class AssignedRolesService(Service):
    """
    Represents a roles sub-collection, for example scoped by user.

    """

    def __init__(self, connection, path):
        super(AssignedRolesService, self).__init__(connection, path)
        self._role_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of roles to return. If not specified all the roles are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def role_service(self, id):
        """
        Sub-resource locator method, returns individual role resource on which the remainder of the URI is dispatched.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return RoleService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.role_service(path)
        return self.role_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AssignedRolesService:%s' % self._path


class AssignedTagService(Service):
    """
    A service to manage assignment of specific tag to specific entities in system.

    """

    def __init__(self, connection, path):
        super(AssignedTagService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets the information about the assigned tag.
        For example to retrieve the information about the tag with the id `456` which is assigned to virtual machine
        with id `123` send a request like this:
        ....
        GET /ovirt-engine/api/vms/123/tags/456
        ....
        [source,xml]
        ----
        <tag href="/ovirt-engine/api/tags/456" id="456">
          <name>root</name>
          <description>root</description>
          <vm href="/ovirt-engine/api/vms/123" id="123"/>
        </tag>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Unassign tag from specific entity in the system.
        For example to unassign the tag with id `456` from virtual machine with id `123` send a request like this:
        ....
        DELETE /ovirt-engine/api/vms/123/tags/456
        ....


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AssignedTagService:%s' % self._path


class AssignedTagsService(Service):
    """
    A service to manage collection of assignment of tags to specific entities in system.

    """

    def __init__(self, connection, path):
        super(AssignedTagsService, self).__init__(connection, path)
        self._tag_service = None

    def add(
        self,
        tag,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Assign tag to specific entity in the system.
        For example to assign tag `mytag` to virtual machine with the id `123` send a request like this:
        ....
        POST /ovirt-engine/api/vms/123/tags
        ....
        With a request body like this:
        [source,xml]
        ----
        <tag>
          <name>mytag</name>
        </tag>
        ----


        This method supports the following parameters:

        `tag`:: The assigned tag.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('tag', tag, types.Tag),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(tag, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List all tags assigned to the specific entity.
        For example to list all the tags of the virtual machine with id `123` send a request like this:
        ....
        GET /ovirt-engine/api/vms/123/tags
        ....
        [source,xml]
        ----
        <tags>
          <tag href="/ovirt-engine/api/tags/222" id="222">
            <name>mytag</name>
            <description>mytag</description>
            <vm href="/ovirt-engine/api/vms/123" id="123"/>
          </tag>
        </tags>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of tags to return. If not specified all the tags are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def tag_service(self, id):
        """
        Reference to the service that manages assignment of specific tag.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return AssignedTagService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.tag_service(path)
        return self.tag_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AssignedTagsService:%s' % self._path


class AssignedVnicProfileService(Service):
    """
    """

    def __init__(self, connection, path):
        super(AssignedVnicProfileService, self).__init__(connection, path)
        self._permissions_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def permissions_service(self):
        """
        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AssignedVnicProfileService:%s' % self._path


class AssignedVnicProfilesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(AssignedVnicProfilesService, self).__init__(connection, path)
        self._profile_service = None

    def add(
        self,
        profile,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('profile', profile, types.VnicProfile),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(profile, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def profile_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return AssignedVnicProfileService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.profile_service(path)
        return self.profile_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AssignedVnicProfilesService:%s' % self._path


class AttachedStorageDomainService(Service):
    """
    """

    def __init__(self, connection, path):
        super(AttachedStorageDomainService, self).__init__(connection, path)
        self._disks_service = None

    def activate(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation activates an attached storage domain.
        Once the storage domain is activated it is ready for use with the data center.
        [source]
        ----
        POST /ovirt-engine/api/datacenters/123/storagedomains/456/activate
        ----
        The activate action does not take any action specific parameters,
        so the request body should contain an empty `action`:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the activation should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'activate', None, headers, query, wait)

    def deactivate(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation deactivates an attached storage domain.
        Once the storage domain is deactivated it will not be used with the data center.
        [source]
        ----
        POST /ovirt-engine/api/datacenters/123/storagedomains/456/deactivate
        ----
        The deactivate action does not take any action specific parameters,
        so the request body should contain an empty `action`:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the deactivation should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'deactivate', None, headers, query, wait)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def disks_service(self):
        """
        """
        return AttachedStorageDomainDisksService(self._connection, '%s/disks' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'disks':
            return self.disks_service()
        if path.startswith('disks/'):
            return self.disks_service().service(path[6:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AttachedStorageDomainService:%s' % self._path


class AttachedStorageDomainDisksService(Service):
    """
    Manages the collection of disks available inside an storage domain that is attached to a data center.

    """

    def __init__(self, connection, path):
        super(AttachedStorageDomainDisksService, self).__init__(connection, path)
        self._disk_service = None

    def add(
        self,
        disk,
        unregistered=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adds or registers a disk.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To add a new disk use the <<services/disks/methods/add, add>>
        operation of the service that manages the disks of the system. To register an unregistered disk use the
        <<services/attached_storage_domain_disk/methods/register, register>> operation of the service that manages
        that disk.


        This method supports the following parameters:

        `disk`:: The disk to add or register.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('disk', disk, types.Disk),
            ('unregistered', unregistered, bool),
        ])

        # Build the URL:
        query = query or {}
        if unregistered is not None:
            unregistered = Writer.render_boolean(unregistered)
            query['unregistered'] = unregistered

        # Send the request and wait for the response:
        return self._internal_add(disk, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieve the list of disks that are available in the storage domain.


        This method supports the following parameters:

        `max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def disk_service(self, id):
        """
        Reference to the service that manages a specific disk.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return AttachedStorageDomainDiskService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.disk_service(path)
        return self.disk_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AttachedStorageDomainDisksService:%s' % self._path


class AttachedStorageDomainsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(AttachedStorageDomainsService, self).__init__(connection, path)
        self._storage_domain_service = None

    def add(
        self,
        storage_domain,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(storage_domain, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of storage domains to return. If not specified all the storage domains are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def storage_domain_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return AttachedStorageDomainService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.storage_domain_service(path)
        return self.storage_domain_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'AttachedStorageDomainsService:%s' % self._path


class BalanceService(Service):
    """
    """

    def __init__(self, connection, path):
        super(BalanceService, self).__init__(connection, path)

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'BalanceService:%s' % self._path


class BalancesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(BalancesService, self).__init__(connection, path)
        self._balance_service = None

    def add(
        self,
        balance,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('balance', balance, types.Balance),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(balance, headers, query, wait)

    def list(
        self,
        filter=None,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of balances to return. If not specified all the balances are returned.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def balance_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return BalanceService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.balance_service(path)
        return self.balance_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'BalancesService:%s' % self._path


class BookmarkService(Service):
    """
    A service to manage a bookmark.

    """

    def __init__(self, connection, path):
        super(BookmarkService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get a bookmark.
        An example for getting a bookmark:
        [source]
        ----
        GET /ovirt-engine/api/bookmarks/123
        ----
        [source,xml]
        ----
        <bookmark href="/ovirt-engine/api/bookmarks/123" id="123">
          <name>example_vm</name>
          <value>vm: name=example*</value>
        </bookmark>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove a bookmark.
        An example for removing a bookmark:
        [source]
        ----
        DELETE /ovirt-engine/api/bookmarks/123
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        bookmark,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Update a bookmark.
        An example for updating a bookmark:
        [source]
        ----
        PUT /ovirt-engine/api/bookmarks/123
        ----
        With the request body:
        [source,xml]
        ----
        <bookmark>
          <name>new_example_vm</name>
          <value>vm: name=new_example*</value>
        </bookmark>
        ----


        This method supports the following parameters:

        `bookmark`:: The updated bookmark.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('bookmark', bookmark, types.Bookmark),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(bookmark, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'BookmarkService:%s' % self._path


class BookmarksService(Service):
    """
    A service to manage bookmarks.

    """

    def __init__(self, connection, path):
        super(BookmarksService, self).__init__(connection, path)
        self._bookmark_service = None

    def add(
        self,
        bookmark,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adding a new bookmark.
        Example of adding a bookmark:
        [source]
        ----
        POST /ovirt-engine/api/bookmarks
        ----
        [source,xml]
        ----
        <bookmark>
          <name>new_example_vm</name>
          <value>vm: name=new_example*</value>
        </bookmark>
        ----


        This method supports the following parameters:

        `bookmark`:: The added bookmark.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('bookmark', bookmark, types.Bookmark),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(bookmark, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Listing all the available bookmarks.
        Example of listing bookmarks:
        [source]
        ----
        GET /ovirt-engine/api/bookmarks
        ----
        [source,xml]
        ----
        <bookmarks>
          <bookmark href="/ovirt-engine/api/bookmarks/123" id="123">
            <name>database</name>
            <value>vm: name=database*</value>
          </bookmark>
          <bookmark href="/ovirt-engine/api/bookmarks/456" id="456">
            <name>example</name>
            <value>vm: name=example*</value>
          </bookmark>
        </bookmarks>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of bookmarks to return. If not specified all the bookmarks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def bookmark_service(self, id):
        """
        A reference to the service managing a specific bookmark.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return BookmarkService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.bookmark_service(path)
        return self.bookmark_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'BookmarksService:%s' % self._path


class ClusterService(Service):
    """
    A service to manage specific cluster.

    """

    def __init__(self, connection, path):
        super(ClusterService, self).__init__(connection, path)
        self._affinity_groups_service = None
        self._cpu_profiles_service = None
        self._gluster_hooks_service = None
        self._gluster_volumes_service = None
        self._network_filters_service = None
        self._networks_service = None
        self._permissions_service = None

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get information about the cluster.
        An example of getting a cluster:
        [source]
        ----
        GET /ovirt-engine/api/clusters/123
        ----
        [source,xml]
        ----
        <cluster href="/ovirt-engine/api/clusters/123" id="123">
          <actions>
            <link href="/ovirt-engine/api/clusters/123/resetemulatedmachine" rel="resetemulatedmachine"/>
          </actions>
          <name>Default</name>
          <description>The default server cluster</description>
          <link href="/ovirt-engine/api/clusters/123/networks" rel="networks"/>
          <link href="/ovirt-engine/api/clusters/123/permissions" rel="permissions"/>
          <link href="/ovirt-engine/api/clusters/123/glustervolumes" rel="glustervolumes"/>
          <link href="/ovirt-engine/api/clusters/123/glusterhooks" rel="glusterhooks"/>
          <link href="/ovirt-engine/api/clusters/123/affinitygroups" rel="affinitygroups"/>
          <link href="/ovirt-engine/api/clusters/123/cpuprofiles" rel="cpuprofiles"/>
          <ballooning_enabled>false</ballooning_enabled>
          <cpu>
            <architecture>x86_64</architecture>
            <type>Intel Penryn Family</type>
          </cpu>
          <error_handling>
            <on_error>migrate</on_error>
          </error_handling>
          <fencing_policy>
            <enabled>true</enabled>
            <skip_if_connectivity_broken>
              <enabled>false</enabled>
              <threshold>50</threshold>
            </skip_if_connectivity_broken>
            <skip_if_sd_active>
              <enabled>false</enabled>
            </skip_if_sd_active>
          </fencing_policy>
          <gluster_service>false</gluster_service>
          <ha_reservation>false</ha_reservation>
          <ksm>
            <enabled>true</enabled>
            <merge_across_nodes>true</merge_across_nodes>
          </ksm>
          <maintenance_reason_required>false</maintenance_reason_required>
          <memory_policy>
            <over_commit>
              <percent>100</percent>
            </over_commit>
            <transparent_hugepages>
              <enabled>true</enabled>
            </transparent_hugepages>
          </memory_policy>
          <migration>
            <auto_converge>inherit</auto_converge>
            <bandwidth>
              <assignment_method>auto</assignment_method>
            </bandwidth>
            <compressed>inherit</compressed>
          </migration>
          <optional_reason>false</optional_reason>
          <required_rng_sources>
            <required_rng_source>random</required_rng_source>
          </required_rng_sources>
          <scheduling_policy href="/ovirt-engine/api/schedulingpolicies/456" id="456"/>
          <threads_as_cores>false</threads_as_cores>
          <trusted_service>false</trusted_service>
          <tunnel_migration>false</tunnel_migration>
          <version>
            <major>4</major>
            <minor>0</minor>
          </version>
          <virt_service>true</virt_service>
          <data_center href="/ovirt-engine/api/datacenters/111" id="111"/>
        </cluster>
        ----


        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes cluster from the system.
        [source]
        ----
        DELETE /ovirt-engine/api/clusters/00000000-0000-0000-0000-000000000000
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def reset_emulated_machine(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the reset should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'resetemulatedmachine', None, headers, query, wait)

    def update(
        self,
        cluster,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates information about the cluster.
        Only specified fields are updated, others remain unchanged.
        E.g. update cluster's CPU:
        [source]
        ----
        PUT /ovirt-engine/api/clusters/123
        ----
        With request body like:
        [source,xml]
        ----
        <cluster>
          <cpu>
            <type>Intel Haswell-noTSX Family</type>
          </cpu>
        </cluster>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('cluster', cluster, types.Cluster),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(cluster, headers, query, wait)

    def affinity_groups_service(self):
        """
        Reference to the service that manages affinity groups.

        """
        return AffinityGroupsService(self._connection, '%s/affinitygroups' % self._path)

    def cpu_profiles_service(self):
        """
        Reference to the service that manages assigned CPU profiles for cluster.

        """
        return AssignedCpuProfilesService(self._connection, '%s/cpuprofiles' % self._path)

    def gluster_hooks_service(self):
        """
        Reference to the service that manages the Gluster hooks for cluster.

        """
        return GlusterHooksService(self._connection, '%s/glusterhooks' % self._path)

    def gluster_volumes_service(self):
        """
        Reference to the service that manages Gluster volumes for cluster.

        """
        return GlusterVolumesService(self._connection, '%s/glustervolumes' % self._path)

    def network_filters_service(self):
        """
        A sub collection with all the supported network filters for this cluster.

        """
        return NetworkFiltersService(self._connection, '%s/networkfilters' % self._path)

    def networks_service(self):
        """
        Reference to the service that manages assigned networks for cluster.

        """
        return AssignedNetworksService(self._connection, '%s/networks' % self._path)

    def permissions_service(self):
        """
        Reference to permissions.

        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'affinitygroups':
            return self.affinity_groups_service()
        if path.startswith('affinitygroups/'):
            return self.affinity_groups_service().service(path[15:])
        if path == 'cpuprofiles':
            return self.cpu_profiles_service()
        if path.startswith('cpuprofiles/'):
            return self.cpu_profiles_service().service(path[12:])
        if path == 'glusterhooks':
            return self.gluster_hooks_service()
        if path.startswith('glusterhooks/'):
            return self.gluster_hooks_service().service(path[13:])
        if path == 'glustervolumes':
            return self.gluster_volumes_service()
        if path.startswith('glustervolumes/'):
            return self.gluster_volumes_service().service(path[15:])
        if path == 'networkfilters':
            return self.network_filters_service()
        if path.startswith('networkfilters/'):
            return self.network_filters_service().service(path[15:])
        if path == 'networks':
            return self.networks_service()
        if path.startswith('networks/'):
            return self.networks_service().service(path[9:])
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ClusterService:%s' % self._path


class ClusterLevelService(Service):
    """
    Provides information about a specific cluster level. See the <<services/cluster_levels,ClusterLevels>> service for
    more information.

    """

    def __init__(self, connection, path):
        super(ClusterLevelService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Provides the information about the capabilities of the specific cluster level managed by this service.
        For example, to find what CPU types are supported by level 3.6 you can send a request like this:
        [source]
        ----
        GET /ovirt-engine/api/clusterlevels/3.6
        ----
        That will return a <<types/cluster_level, ClusterLevel>> object containing the supported CPU types, and other
        information which describes the cluster level:
        [source,xml]
        ----
        <cluster_level id="3.6">
          <cpu_types>
            <cpu_type>
              <name>Intel Conroe Family</name>
              <level>3</level>
              <architecture>x86_64</architecture>
            </cpu_type>
            ...
          </cpu_types>
          <permits>
            <permit id="1">
              <name>create_vm</name>
              <administrative>false</administrative>
            </permit>
            ...
          </permits>
        </cluster_level>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ClusterLevelService:%s' % self._path


class ClusterLevelsService(Service):
    """
    Provides information about the capabilities of different cluster levels supported by the engine. Version 4.0 of the
    engine supports levels 4.0 and 3.6. Each of these levels support different sets of CPU types, for example. This
    service provides that information.

    """

    def __init__(self, connection, path):
        super(ClusterLevelsService, self).__init__(connection, path)
        self._level_service = None

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists the cluster levels supported by the system.
        [source]
        ----
        GET /ovirt-engine/api/clusterlevels
        ----
        This will return a list of available cluster levels.
        [source,xml]
        ----
        <cluster_levels>
          <cluster_level id="4.0">
             ...
          </cluster_level>
          ...
        </cluster_levels>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def level_service(self, id):
        """
        Reference to the service that provides information about an specific cluster level.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return ClusterLevelService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.level_service(path)
        return self.level_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'ClusterLevelsService:%s' % self._path


class ClustersService(Service):
    """
    A service to manage clusters.

    """

    def __init__(self, connection, path):
        super(ClustersService, self).__init__(connection, path)
        self._cluster_service = None

    def add(
        self,
        cluster,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new cluster.
        This requires the `name`, `cpu.type` and `data_center` attributes. Identify the data center with either the `id`
        or `name` attributes.
        [source]
        ----
        POST /ovirt-engine/api/clusters
        ----
        With a request body like this:
        [source,xml]
        ----
        <cluster>
          <name>mycluster</name>
          <cpu>
            <type>Intel Penryn Family</type>
          </cpu>
          <data_center id="123"/>
        </cluster>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('cluster', cluster, types.Cluster),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(cluster, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        filter=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of clusters to return. If not specified all the clusters are returned.

        `search`:: A query string used to restrict the returned clusters.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('filter', filter, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def cluster_service(self, id):
        """
        Reference to the service that manages a specific cluster.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return ClusterService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.cluster_service(path)
        return self.cluster_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'ClustersService:%s' % self._path


class CopyableService(Service):
    """
    """

    def __init__(self, connection, path):
        super(CopyableService, self).__init__(connection, path)

    def copy(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the copy should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'copy', None, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'CopyableService:%s' % self._path


class CpuProfileService(Service):
    """
    """

    def __init__(self, connection, path):
        super(CpuProfileService, self).__init__(connection, path)
        self._permissions_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        profile,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('profile', profile, types.CpuProfile),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(profile, headers, query, wait)

    def permissions_service(self):
        """
        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'CpuProfileService:%s' % self._path


class CpuProfilesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(CpuProfilesService, self).__init__(connection, path)
        self._profile_service = None

    def add(
        self,
        profile,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('profile', profile, types.CpuProfile),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(profile, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def profile_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return CpuProfileService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.profile_service(path)
        return self.profile_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'CpuProfilesService:%s' % self._path


class DataCenterService(Service):
    """
    A service to manage a data center.

    """

    def __init__(self, connection, path):
        super(DataCenterService, self).__init__(connection, path)
        self._clusters_service = None
        self._iscsi_bonds_service = None
        self._networks_service = None
        self._permissions_service = None
        self._qoss_service = None
        self._quotas_service = None
        self._storage_domains_service = None

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get a data center.
        An example of getting a data center:
        [source]
        ----
        GET /ovirt-engine/api/datacenters/123
        ----
        [source,xml]
        ----
        <data_center href="/ovirt-engine/api/datacenters/123" id="123">
          <name>Default</name>
          <description>The default Data Center</description>
          <link href="/ovirt-engine/api/datacenters/123/clusters" rel="clusters"/>
          <link href="/ovirt-engine/api/datacenters/123/storagedomains" rel="storagedomains"/>
          <link href="/ovirt-engine/api/datacenters/123/permissions" rel="permissions"/>
          <link href="/ovirt-engine/api/datacenters/123/networks" rel="networks"/>
          <link href="/ovirt-engine/api/datacenters/123/quotas" rel="quotas"/>
          <link href="/ovirt-engine/api/datacenters/123/qoss" rel="qoss"/>
          <link href="/ovirt-engine/api/datacenters/123/iscsibonds" rel="iscsibonds"/>
          <local>false</local>
          <quota_mode>disabled</quota_mode>
          <status>up</status>
          <storage_format>v3</storage_format>
          <supported_versions>
            <version>
              <major>4</major>
              <minor>0</minor>
           </version>
          </supported_versions>
          <version>
            <major>4</major>
            <minor>0</minor>
          </version>
          <mac_pool href="/ovirt-engine/api/macpools/456" id="456"/>
        </data_center>
        ----


        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        force=None,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the data center.
        [source]
        ----
        DELETE /ovirt-engine/api/datacenters/123
        ----
        Without any special parameters, the storage domains attached to the data center are detached and then removed
        from the storage. If something fails when performing this operation, for example if there is no host available to
        remove the storage domains from the storage, the complete operation will fail.
        If the `force` parameter is `true` then the operation will always succeed, even if something fails while removing
        one storage domain, for example. The failure is just ignored and the data center is removed from the database
        anyway.


        This method supports the following parameters:

        `force`:: Indicates if the operation should succeed, and the storage domain removed from the database, even if
        something fails during the operation.
        This parameter is optional, and the default value is `false`.

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('force', force, bool),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if force is not None:
            force = Writer.render_boolean(force)
            query['force'] = force
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        data_center,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the data center.
        The `name`, `description`, `storage_type`, `version`, `storage_format` and `mac_pool` elements are updatable
        post-creation. For example, to change the name and description of data center `123` send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/datacenters/123
        ----
        With a request body like this:
        [source,xml]
        ----
        <data_center>
          <name>myupdatedname</name>
          <description>An updated description for the data center</description>
        </data_center>
        ----


        This method supports the following parameters:

        `data_center`:: The data center that is being updated.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('data_center', data_center, types.DataCenter),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(data_center, headers, query, wait)

    def clusters_service(self):
        """
        """
        return ClustersService(self._connection, '%s/clusters' % self._path)

    def iscsi_bonds_service(self):
        """
        Reference to the iSCSI bonds service.

        """
        return IscsiBondsService(self._connection, '%s/iscsibonds' % self._path)

    def networks_service(self):
        """
        Returns a reference to the service, that manages the networks, that are associated with the data center.

        """
        return NetworksService(self._connection, '%s/networks' % self._path)

    def permissions_service(self):
        """
        Reference to the permissions service.

        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def qoss_service(self):
        """
        Reference to the QOSs service.

        """
        return QossService(self._connection, '%s/qoss' % self._path)

    def quotas_service(self):
        """
        Reference to the quotas service.

        """
        return QuotasService(self._connection, '%s/quotas' % self._path)

    def storage_domains_service(self):
        """
        Attach and detach storage domains to and from a data center.
        For attaching a single storage domain we should use the following POST request:
        [source]
        ----
        POST /ovirt-engine/api/datacenters/123/storagedomains
        ----
        With a request body like this:
        [source,xml]
        ----
        <storage_domain>
          <name>data1</name>
        </storage_domain>
        ----
        For detaching a single storage domain we should use the following DELETE request:
        [source]
        ----
        DELETE /ovirt-engine/api/datacenters/123/storagedomains/123
        ----

        """
        return AttachedStorageDomainsService(self._connection, '%s/storagedomains' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'clusters':
            return self.clusters_service()
        if path.startswith('clusters/'):
            return self.clusters_service().service(path[9:])
        if path == 'iscsibonds':
            return self.iscsi_bonds_service()
        if path.startswith('iscsibonds/'):
            return self.iscsi_bonds_service().service(path[11:])
        if path == 'networks':
            return self.networks_service()
        if path.startswith('networks/'):
            return self.networks_service().service(path[9:])
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'qoss':
            return self.qoss_service()
        if path.startswith('qoss/'):
            return self.qoss_service().service(path[5:])
        if path == 'quotas':
            return self.quotas_service()
        if path.startswith('quotas/'):
            return self.quotas_service().service(path[7:])
        if path == 'storagedomains':
            return self.storage_domains_service()
        if path.startswith('storagedomains/'):
            return self.storage_domains_service().service(path[15:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'DataCenterService:%s' % self._path


class DataCentersService(Service):
    """
    A service to manage data centers.

    """

    def __init__(self, connection, path):
        super(DataCentersService, self).__init__(connection, path)
        self._data_center_service = None

    def add(
        self,
        data_center,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new data center.
        Creation of a new data center requires the `name` and `local` elements. For example, to create a data center
        named `mydc` that uses shared storage (NFS, iSCSI or fibre channel) send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/datacenters
        ----
        With a request body like this:
        [source,xml]
        ----
        <data_center>
          <name>mydc</name>
          <local>false</local>
        </data_center>
        ----


        This method supports the following parameters:

        `data_center`:: The data center that is being added.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('data_center', data_center, types.DataCenter),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(data_center, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        filter=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists the data centers.
        The following request retrieves a representation of the data centers:
        [source]
        ----
        GET /ovirt-engine/api/datacenters
        ----
        The above request performed with `curl`:
        [source,bash]
        ----
        curl \
        --request GET \
        --cacert /etc/pki/ovirt-engine/ca.pem \
        --header "Version: 4" \
        --header "Accept: application/xml" \
        --user "admin@internal:mypassword" \
        https://myengine.example.com/ovirt-engine/api/datacenters
        ----
        This is what an example response could look like:
        [source,xml]
        ----
        <data_center href="/ovirt-engine/api/datacenters/123" id="123">
          <name>Default</name>
          <description>The default Data Center</description>
          <link href="/ovirt-engine/api/datacenters/123/networks" rel="networks"/>
          <link href="/ovirt-engine/api/datacenters/123/storagedomains" rel="storagedomains"/>
          <link href="/ovirt-engine/api/datacenters/123/permissions" rel="permissions"/>
          <link href="/ovirt-engine/api/datacenters/123/clusters" rel="clusters"/>
          <link href="/ovirt-engine/api/datacenters/123/qoss" rel="qoss"/>
          <link href="/ovirt-engine/api/datacenters/123/iscsibonds" rel="iscsibonds"/>
          <link href="/ovirt-engine/api/datacenters/123/quotas" rel="quotas"/>
          <local>false</local>
          <quota_mode>disabled</quota_mode>
          <status>up</status>
          <supported_versions>
            <version>
              <major>4</major>
              <minor>0</minor>
            </version>
          </supported_versions>
          <version>
            <major>4</major>
            <minor>0</minor>
          </version>
        </data_center>
        ----
        Note the `id` code of your `Default` data center. This code identifies this data center in relation to other
        resources of your virtual environment.
        The data center also contains a link to the storage domains collection. The data center uses this collection to
        attach storage domains from the storage domains main collection.


        This method supports the following parameters:

        `max`:: Sets the maximum number of data centers to return. If not specified all the data centers are returned.

        `search`:: A query string used to restrict the returned data centers.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('filter', filter, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def data_center_service(self, id):
        """
        Reference to the service that manages a specific data center.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return DataCenterService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.data_center_service(path)
        return self.data_center_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'DataCentersService:%s' % self._path


class DiskAttachmentService(Service):
    """
    This service manages the attachment of a disk to a virtual machine.

    """

    def __init__(self, connection, path):
        super(DiskAttachmentService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns the details of the attachment, including the bootable flag and link to the disk.
        An example of getting a disk attachment:
        [source]
        ----
        GET /ovirt-engine/api/vms/123/diskattachments/456
        ----
        [source,xml]
        ----
        <disk_attachment href="/ovirt-engine/api/vms/123/diskattachments/456" id="456">
          <active>true</active>
          <bootable>true</bootable>
          <interface>virtio</interface>
          <disk href="/ovirt-engine/api/disks/456" id="456"/>
          <vm href="/ovirt-engine/api/vms/123" id="123"/>
        </disk_attachment>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        detach_only=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the disk attachment.
        This will only detach the disk from the virtual machine, but won't remove it from
        the system, unless the `detach_only` parameter is `false`.
        An example of removing a disk attachment:
        [source]
        ----
        DELETE /ovirt-engine/api/vms/123/diskattachments/456?detach_only=true
        ----


        This method supports the following parameters:

        `detach_only`:: Indicates if the disk should only be detached from the virtual machine, but not removed from the system.
        The default value is `true`, which won't remove the disk from the system.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('detach_only', detach_only, bool),
        ])

        # Build the URL:
        query = query or {}
        if detach_only is not None:
            detach_only = Writer.render_boolean(detach_only)
            query['detach_only'] = detach_only

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        disk_attachment,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Update the disk attachment and the disk properties within it.
        [source]
        ----
        PUT /vms/{vm:id}/disksattachments/{attachment:id}
        <disk_attachment>
          <bootable>true</bootable>
          <interface>ide</interface>
          <active>true</active>
          <disk>
            <name>mydisk</name>
            <provisioned_size>1024</provisioned_size>
            ...
          </disk>
        </disk_attachment>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('disk_attachment', disk_attachment, types.DiskAttachment),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_update(disk_attachment, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'DiskAttachmentService:%s' % self._path


class DiskAttachmentsService(Service):
    """
    This service manages the set of disks attached to a virtual machine. Each attached disk is represented by a
    <<types/disk_attachment,DiskAttachment>>, containing the bootable flag, the disk interface and the reference to
    the disk.

    """

    def __init__(self, connection, path):
        super(DiskAttachmentsService, self).__init__(connection, path)
        self._attachment_service = None

    def add(
        self,
        attachment,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adds a new disk attachment to the virtual machine. The `attachment` parameter can contain just a reference, if
        the disk already exists:
        [source,xml]
        ----
        <disk_attachment>
          <bootable>true</bootable>
          <pass_discard>true</pass_discard>
          <interface>ide</interface>
          <active>true</active>
          <disk id="123"/>
        </disk_attachment>
        ----
        Or it can contain the complete representation of the disk, if the disk doesn't exist yet:
        [source,xml]
        ----
        <disk_attachment>
          <bootable>true</bootable>
          <pass_discard>true</pass_discard>
          <interface>ide</interface>
          <active>true</active>
          <disk>
            <name>mydisk</name>
            <provisioned_size>1024</provisioned_size>
            ...
          </disk>
        </disk_attachment>
        ----
        In this case the disk will be created and then attached to the virtual machine.
        In both cases, use the following URL for a virtual machine with an id `345`:
        [source]
        ----
        POST /ovirt-engine/api/vms/345/diskattachments
        ----
        IMPORTANT: The server accepts requests that don't contain the `active` attribute, but the effect is
        undefined. In some cases the disk will be automatically activated and in other cases it won't. To
        avoid issues it is strongly recommended to always include the `active` attribute with the desired
        value.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('attachment', attachment, types.DiskAttachment),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(attachment, headers, query, wait)

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List the disk that are attached to the virtual machine.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def attachment_service(self, id):
        """
        Reference to the service that manages a specific attachment.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return DiskAttachmentService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.attachment_service(path)
        return self.attachment_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'DiskAttachmentsService:%s' % self._path


class DiskProfileService(Service):
    """
    """

    def __init__(self, connection, path):
        super(DiskProfileService, self).__init__(connection, path)
        self._permissions_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        profile,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('profile', profile, types.DiskProfile),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(profile, headers, query, wait)

    def permissions_service(self):
        """
        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'DiskProfileService:%s' % self._path


class DiskProfilesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(DiskProfilesService, self).__init__(connection, path)
        self._disk_profile_service = None

    def add(
        self,
        profile,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('profile', profile, types.DiskProfile),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(profile, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def disk_profile_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return DiskProfileService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.disk_profile_service(path)
        return self.disk_profile_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'DiskProfilesService:%s' % self._path


class DiskSnapshotService(Service):
    """
    """

    def __init__(self, connection, path):
        super(DiskSnapshotService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'DiskSnapshotService:%s' % self._path


class DiskSnapshotsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(DiskSnapshotsService, self).__init__(connection, path)
        self._snapshot_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of snapshots to return. If not specified all the snapshots are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def snapshot_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return DiskSnapshotService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.snapshot_service(path)
        return self.snapshot_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'DiskSnapshotsService:%s' % self._path


class DisksService(Service):
    """
    Manages the collection of disks available in the system.

    """

    def __init__(self, connection, path):
        super(DisksService, self).__init__(connection, path)
        self._disk_service = None

    def add(
        self,
        disk,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adds a new floating disk.
        There are three types of disks that can be added - disk image, direct LUN and
         https://wiki.openstack.org/wiki/Cinder[Cinder] disk.
        *Adding a new image disk:*
        When creating a new floating image <<types/disk,Disk>>, the API requires the `storage_domain`, `provisioned_size`
        and `format` attributes.
        To create a new floating image disk with specified `provisioned_size`, `format` and `name` on a storage domain
        with an id `123`, send a request as follows:
        [source]
        ----
        POST /ovirt-engine/api/disks
        ----
        With a request body as follows:
        [source,xml]
        ----
        <disk>
          <storage_domains>
            <storage_domain id="123"/>
          </storage_domains>
          <name>mydisk</name>
          <provisioned_size>1048576</provisioned_size>
          <format>cow</format>
        </disk>
        ----
        *Adding a new direct LUN disk:*
        When adding a new floating direct LUN via the API, there are two flavors that can be used:
        . With a `host` element - in this case, the host is used for sanity checks (e.g., that the LUN is visible) and
        to retrieve basic information about the LUN (e.g., size and serial).
        . Without a `host` element - in this case, the operation is a database-only operation, and the storage is never
        accessed.
        To create a new floating direct LUN disk with a `host` element with an id `123`, specified `alias`, `type` and
        `logical_unit` with an id `456` (that has the attributes `address`, `port` and `target`),
        send a request as follows:
        [source]
        ----
        POST /ovirt-engine/api/disks
        ----
        With a request body as follows:
        [source,xml]
        ----
        <disk>
          <alias>mylun</alias>
          <lun_storage>
            <host id="123"/>
            <type>iscsi</type>
            <logical_units>
              <logical_unit id="456">
                <address>10.35.10.20</address>
                <port>3260</port>
                <target>iqn.2017-01.com.myhost:444</target>
              </logical_unit>
            </logical_units>
          </lun_storage>
        </disk>
        ----
        To create a new floating direct LUN disk without using a host, remove the `host` element.
        *Adding a new Cinder disk:*
        To create a new floating Cinder disk, send a request as follows:
        [source]
        ----
        POST /ovirt-engine/api/disks
        ----
        With a request body as follows:
        [source,xml]
        ----
        <disk>
          <openstack_volume_type>
            <name>myceph</name>
          </openstack_volume_type>
          <storage_domains>
            <storage_domain>
              <name>cinderDomain</name>
            </storage_domain>
          </storage_domains>
          <provisioned_size>1073741824</provisioned_size>
          <interface>virtio</interface>
          <format>raw</format>
        </disk>
        ----


        This method supports the following parameters:

        `disk`:: The disk.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('disk', disk, types.Disk),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(disk, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get list of disks.
        [source]
        ----
        GET /ovirt-engine/api/disks
        ----
        You will get a XML response which will look like this one:
        [source,xml]
        ----
        <disks>
          <disk id="123">
            <actions>...</actions>
            <name>MyDisk</name>
            <description>MyDisk description</description>
            <link href="/ovirt-engine/api/disks/123/permissions" rel="permissions"/>
            <link href="/ovirt-engine/api/disks/123/statistics" rel="statistics"/>
            <actual_size>5345845248</actual_size>
            <alias>MyDisk alias</alias>
            ...
            <status>ok</status>
            <storage_type>image</storage_type>
            <wipe_after_delete>false</wipe_after_delete>
            <disk_profile id="123"/>
            <quota id="123"/>
            <storage_domains>...</storage_domains>
          </disk>
          ...
        </disks>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.

        `search`:: A query string used to restrict the returned disks.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def disk_service(self, id):
        """
        Reference to a service managing a specific disk.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return DiskService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.disk_service(path)
        return self.disk_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'DisksService:%s' % self._path


class DomainService(Service):
    """
    A service to view details of an authentication domain in the system.

    """

    def __init__(self, connection, path):
        super(DomainService, self).__init__(connection, path)
        self._groups_service = None
        self._users_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets the authentication domain information.
        Usage:
        ....
        GET /ovirt-engine/api/domains/5678
        ....
        Will return the domain information:
        [source,xml]
        ----
        <domain href="/ovirt-engine/api/domains/5678" id="5678">
          <name>internal-authz</name>
          <link href="/ovirt-engine/api/domains/5678/users" rel="users"/>
          <link href="/ovirt-engine/api/domains/5678/groups" rel="groups"/>
          <link href="/ovirt-engine/api/domains/5678/users?search={query}" rel="users/search"/>
          <link href="/ovirt-engine/api/domains/5678/groups?search={query}" rel="groups/search"/>
        </domain>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def groups_service(self):
        """
        Reference to a service to manage domain groups.

        """
        return DomainGroupsService(self._connection, '%s/groups' % self._path)

    def users_service(self):
        """
        Reference to a service to manage domain users.

        """
        return DomainUsersService(self._connection, '%s/users' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'groups':
            return self.groups_service()
        if path.startswith('groups/'):
            return self.groups_service().service(path[7:])
        if path == 'users':
            return self.users_service()
        if path.startswith('users/'):
            return self.users_service().service(path[6:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'DomainService:%s' % self._path


class DomainGroupService(Service):
    """
    """

    def __init__(self, connection, path):
        super(DomainGroupService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'DomainGroupService:%s' % self._path


class DomainGroupsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(DomainGroupsService, self).__init__(connection, path)
        self._group_service = None

    def list(
        self,
        case_sensitive=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of groups to return. If not specified all the groups are returned.

        `search`:: A query string used to restrict the returned groups.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def group_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return DomainGroupService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.group_service(path)
        return self.group_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'DomainGroupsService:%s' % self._path


class DomainUserService(Service):
    """
    A service to view a domain user in the system.

    """

    def __init__(self, connection, path):
        super(DomainUserService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets the domain user information.
        Usage:
        ....
        GET /ovirt-engine/api/domains/5678/users/1234
        ....
        Will return the domain user information:
        [source,xml]
        ----
        <user href="/ovirt-engine/api/users/1234" id="1234">
          <name>admin</name>
          <namespace>*</namespace>
          <principal>admin</principal>
          <user_name>admin@internal-authz</user_name>
          <domain href="/ovirt-engine/api/domains/5678" id="5678">
            <name>internal-authz</name>
          </domain>
          <groups/>
        </user>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'DomainUserService:%s' % self._path


class DomainUsersService(Service):
    """
    A service to list all domain users in the system.

    """

    def __init__(self, connection, path):
        super(DomainUsersService, self).__init__(connection, path)
        self._user_service = None

    def list(
        self,
        case_sensitive=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List all the users in the domain.
        Usage:
        ....
        GET /ovirt-engine/api/domains/5678/users
        ....
        Will return the list of users in the domain:
        [source,xml]
        ----
        <users>
          <user href="/ovirt-engine/api/domains/5678/users/1234" id="1234">
            <name>admin</name>
            <namespace>*</namespace>
            <principal>admin</principal>
            <user_name>admin@internal-authz</user_name>
            <domain href="/ovirt-engine/api/domains/5678" id="5678">
              <name>internal-authz</name>
            </domain>
            <groups/>
          </user>
        </users>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of users to return. If not specified all the users are returned.

        `search`:: A query string used to restrict the returned users.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def user_service(self, id):
        """
        Reference to a service to view details of a domain user.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return DomainUserService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.user_service(path)
        return self.user_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'DomainUsersService:%s' % self._path


class DomainsService(Service):
    """
    A service to list all authentication domains in the system.

    """

    def __init__(self, connection, path):
        super(DomainsService, self).__init__(connection, path)
        self._domain_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List all the authentication domains in the system.
        Usage:
        ....
        GET /ovirt-engine/api/domains
        ....
        Will return the list of domains:
        [source,xml]
        ----
        <domains>
          <domain href="/ovirt-engine/api/domains/5678" id="5678">
            <name>internal-authz</name>
            <link href="/ovirt-engine/api/domains/5678/users" rel="users"/>
            <link href="/ovirt-engine/api/domains/5678/groups" rel="groups"/>
            <link href="/ovirt-engine/api/domains/5678/users?search={query}" rel="users/search"/>
            <link href="/ovirt-engine/api/domains/5678/groups?search={query}" rel="groups/search"/>
          </domain>
        </domains>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of domains to return. If not specified all the domains are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def domain_service(self, id):
        """
        Reference to a service to view details of a domain.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return DomainService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.domain_service(path)
        return self.domain_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'DomainsService:%s' % self._path


class EventService(Service):
    """
    A service to manage an event in the system.

    """

    def __init__(self, connection, path):
        super(EventService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get an event.
        An example of getting an event:
        [source]
        ----
        GET /ovirt-engine/api/events/123
        ----
        [source,xml]
        ----
        <event href="/ovirt-engine/api/events/123" id="123">
          <description>Host example.com was added by admin@internal-authz.</description>
          <code>42</code>
          <correlation_id>135</correlation_id>
          <custom_id>-1</custom_id>
          <flood_rate>30</flood_rate>
          <origin>oVirt</origin>
          <severity>normal</severity>
          <time>2016-12-11T11:13:44.654+02:00</time>
          <cluster href="/ovirt-engine/api/clusters/456" id="456"/>
          <host href="/ovirt-engine/api/hosts/789" id="789"/>
          <user href="/ovirt-engine/api/users/987" id="987"/>
        </event>
        ----
        Note that the number of fields changes according to the information that resides on the event.
        For example, for storage domain related events you will get the storage domain reference,
        as well as the reference for the data center this storage domain resides in.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes an event from internal audit log.
        An event can be removed by sending following request
        [source]
        ----
        DELETE /ovirt-engine/api/events/123
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'EventService:%s' % self._path


class EventsService(Service):
    """
    A service to manage events in the system.

    """

    def __init__(self, connection, path):
        super(EventsService, self).__init__(connection, path)
        self._event_service = None

    def add(
        self,
        event,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adds an external event to the internal audit log.
        This is intended for integration with external systems that detect or produce events relevant for the
        administrator of the system. For example, an external monitoring tool may be able to detect that a file system
        is full inside the guest operating system of a virtual machine. This event can be added to the internal audit
        log sending a request like this:
        [source]
        ----
        POST /ovirt-engine/api/events
        <event>
          <description>File system /home is full</description>
          <severity>alert</severity>
          <origin>mymonitor</origin>
          <custom_id>1467879754</custom_id>
        </event>
        ----
        Events can also be linked to specific objects. For example, the above event could be linked to the specific
        virtual machine where it happened, using the `vm` link:
        [source]
        ----
        POST /ovirt-engine/api/events
        <event>
          <description>File system /home is full</description>
          <severity>alert</severity>
          <origin>mymonitor</origin>
          <custom_id>1467879754</custom_id>
          <vm id="aae98225-5b73-490d-a252-899209af17e9"/>
        </event>
        ----
        NOTE: When using links, like the `vm` in the previous example, only the `id` attribute is accepted. The `name`
        attribute, if provided, is simply ignored.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('event', event, types.Event),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(event, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        from_=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get list of events.
        [source]
        ----
        GET /ovirt-engine/api/events
        ----
        To the above request we get following response:
        [source,xml]
        ----
        <events>
          <event href="/ovirt-engine/api/events/2" id="2">
            <description>User admin@internal-authz logged out.</description>
            <code>31</code>
            <correlation_id>1e892ea9</correlation_id>
            <custom_id>-1</custom_id>
            <flood_rate>30</flood_rate>
            <origin>oVirt</origin>
            <severity>normal</severity>
            <time>2016-09-14T12:14:34.541+02:00</time>
            <user href="/ovirt-engine/api/users/57d91d48-00da-0137-0138-000000000244" id="57d91d48-00da-0137-0138-000000000244"/>
          </event>
          <event href="/ovirt-engine/api/events/1" id="1">
            <description>User admin logged in.</description>
            <code>30</code>
            <correlation_id>1fbd81f4</correlation_id>
            <custom_id>-1</custom_id>
            <flood_rate>30</flood_rate>
            <origin>oVirt</origin>
            <severity>normal</severity>
            <time>2016-09-14T11:54:35.229+02:00</time>
            <user href="/ovirt-engine/api/users/57d91d48-00da-0137-0138-000000000244" id="57d91d48-00da-0137-0138-000000000244"/>
          </event>
        </events>
        ----
        The following events occur:
        * id="1" - The API logs in the admin user account.
        * id="2" - The API logs out of the admin user account.


        This method supports the following parameters:

        `from_`:: Indicates the identifier of the the first event that should be returned. The identifiers of events are
        strictly increasing, so when this parameter is used only the events with that identifiers equal or greater
        than the given value will be returned. For example, the following request will return only the events
        with identifiers greater or equal than `123`:
        [source]
        ----
        GET /ovirt-engine/api/events?from=123
        ----
        This parameter is optional, and if not specified then the first event returned will be most recently
        generated.

        `max`:: Sets the maximum number of events to return. If not specified all the events are returned.

        `search`:: The events service provides search queries similar to other resource services.
        We can search by providing specific severity.
        [source]
        ----
        GET /ovirt-engine/api/events?search=severity%3Dnormal
        ----
        To the above request we get a list of events which severity is equal to `normal`:
        [source,xml]
        ----
        <events>
          <event href="/ovirt-engine/api/events/2" id="2">
            <description>User admin@internal-authz logged out.</description>
            <code>31</code>
            <correlation_id>1fbd81f4</correlation_id>
            <custom_id>-1</custom_id>
            <flood_rate>30</flood_rate>
            <origin>oVirt</origin>
            <severity>normal</severity>
            <time>2016-09-14T11:54:35.229+02:00</time>
            <user href="/ovirt-engine/api/users/57d91d48-00da-0137-0138-000000000244" id="57d91d48-00da-0137-0138-000000000244"/>
          </event>
          <event href="/ovirt-engine/api/events/1" id="1">
            <description>Affinity Rules Enforcement Manager started.</description>
            <code>10780</code>
            <custom_id>-1</custom_id>
            <flood_rate>30</flood_rate>
            <origin>oVirt</origin>
            <severity>normal</severity>
            <time>2016-09-14T11:52:18.861+02:00</time>
          </event>
        </events>
        ----
        A virtualization environment generates a large amount of events after
        a period of time. However, the API only displays a default number of
        events for one search query. To display more than the default, the API
        separates results into pages with the page command in a search query.
        The following search query tells the API to paginate results using a
        page value in combination with the sortby clause:
        [source]
        ----
        sortby time asc page 1
        ----
        Below example paginates event resources. The URL-encoded request is:
        [source]
        ----
        GET /ovirt-engine/api/events?search=sortby%20time%20asc%20page%201
        ----
        Increase the page value to view the next page of results.
        [source]
        ----
        GET /ovirt-engine/api/events?search=sortby%20time%20asc%20page%202
        ----

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('from_', from_, int),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if from_ is not None:
            from_ = Writer.render_integer(from_)
            query['from'] = from_
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def undelete(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the un-delete should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'undelete', None, headers, query, wait)

    def event_service(self, id):
        """
        Reference to the service that manages a specific event.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return EventService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.event_service(path)
        return self.event_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'EventsService:%s' % self._path


class ExternalComputeResourceService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalComputeResourceService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ExternalComputeResourceService:%s' % self._path


class ExternalComputeResourcesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalComputeResourcesService, self).__init__(connection, path)
        self._resource_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of resources to return. If not specified all the resources are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def resource_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return ExternalComputeResourceService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.resource_service(path)
        return self.resource_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'ExternalComputeResourcesService:%s' % self._path


class ExternalDiscoveredHostService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalDiscoveredHostService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ExternalDiscoveredHostService:%s' % self._path


class ExternalDiscoveredHostsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalDiscoveredHostsService, self).__init__(connection, path)
        self._host_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of hosts to return. If not specified all the hosts are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def host_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return ExternalDiscoveredHostService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.host_service(path)
        return self.host_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'ExternalDiscoveredHostsService:%s' % self._path


class ExternalHostService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalHostService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ExternalHostService:%s' % self._path


class ExternalHostGroupService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalHostGroupService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ExternalHostGroupService:%s' % self._path


class ExternalHostGroupsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalHostGroupsService, self).__init__(connection, path)
        self._group_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of groups to return. If not specified all the groups are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def group_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return ExternalHostGroupService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.group_service(path)
        return self.group_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'ExternalHostGroupsService:%s' % self._path


class ExternalHostProvidersService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalHostProvidersService, self).__init__(connection, path)
        self._provider_service = None

    def add(
        self,
        provider,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('provider', provider, types.ExternalHostProvider),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(provider, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of providers to return. If not specified all the providers are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def provider_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return ExternalHostProviderService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.provider_service(path)
        return self.provider_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'ExternalHostProvidersService:%s' % self._path


class ExternalHostsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalHostsService, self).__init__(connection, path)
        self._host_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of hosts to return. If not specified all the hosts are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def host_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return ExternalHostService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.host_service(path)
        return self.host_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'ExternalHostsService:%s' % self._path


class ExternalProviderService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalProviderService, self).__init__(connection, path)
        self._certificates_service = None

    def import_certificates(
        self,
        certificates=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('certificates', certificates, list),
        ])

        # Populate the action:
        action = types.Action(
            certificates=certificates,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'importcertificates', None, headers, query, wait)

    def test_connectivity(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the test should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'testconnectivity', None, headers, query, wait)

    def certificates_service(self):
        """
        """
        return ExternalProviderCertificatesService(self._connection, '%s/certificates' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'certificates':
            return self.certificates_service()
        if path.startswith('certificates/'):
            return self.certificates_service().service(path[13:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ExternalProviderService:%s' % self._path


class ExternalProviderCertificateService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalProviderCertificateService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ExternalProviderCertificateService:%s' % self._path


class ExternalProviderCertificatesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ExternalProviderCertificatesService, self).__init__(connection, path)
        self._certificate_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of certificates to return. If not specified all the certificates are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def certificate_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return ExternalProviderCertificateService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.certificate_service(path)
        return self.certificate_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'ExternalProviderCertificatesService:%s' % self._path


class ExternalVmImportsService(Service):
    """
    Provides capability to import external virtual machines.

    """

    def __init__(self, connection, path):
        super(ExternalVmImportsService, self).__init__(connection, path)

    def add(
        self,
        import_,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation is used to import a virtual machine from external hypervisor, such as KVM, XEN or VMware.
        For example import of a virtual machine from VMware can be facilitated using the following request:
        [source]
        ----
        POST /externalvmimports
        ----
        With request body of type <<types/external_vm_import,ExternalVmImport>>, for example:
        [source,xml]
        ----
        <external_vm_import>
          <vm>
            <name>my_vm</name>
          </vm>
          <cluster id="360014051136c20574f743bdbd28177fd" />
          <storage_domain id="8bb5ade5-e988-4000-8b93-dbfc6717fe50" />
          <name>vm_name_as_is_in_vmware</name>
          <sparse>true</sparse>
          <username>vmware_user</username>
          <password>123456</password>
          <provider>VMWARE</provider>
          <url>vpx://wmware_user@vcenter-host/DataCenter/Cluster/esxi-host?no_verify=1</url>
          <drivers_iso id="virtio-win-1.6.7.iso" />
        </external_vm_import>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('import_', import_, types.ExternalVmImport),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(import_, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ExternalVmImportsService:%s' % self._path


class FenceAgentService(Service):
    """
    """

    def __init__(self, connection, path):
        super(FenceAgentService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        agent,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('agent', agent, types.Agent),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(agent, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'FenceAgentService:%s' % self._path


class FenceAgentsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(FenceAgentsService, self).__init__(connection, path)
        self._agent_service = None

    def add(
        self,
        agent,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('agent', agent, types.Agent),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(agent, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of agents to return. If not specified all the agents are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def agent_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return FenceAgentService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.agent_service(path)
        return self.agent_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'FenceAgentsService:%s' % self._path


class FileService(Service):
    """
    """

    def __init__(self, connection, path):
        super(FileService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'FileService:%s' % self._path


class FilesService(Service):
    """
    Provides a way for clients to list available files.
    This services is specifically targeted to ISO storage domains, which contain ISO images and virtual floppy disks
    (VFDs) that an administrator uploads.
    The addition of a CDROM device to a virtual machine requires an ISO image from the files of an ISO storage domain.

    """

    def __init__(self, connection, path):
        super(FilesService, self).__init__(connection, path)
        self._file_service = None

    def list(
        self,
        case_sensitive=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of files to return. If not specified all the files are returned.

        `search`:: A query string used to restrict the returned files.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def file_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return FileService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.file_service(path)
        return self.file_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'FilesService:%s' % self._path


class FilterService(Service):
    """
    """

    def __init__(self, connection, path):
        super(FilterService, self).__init__(connection, path)

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'FilterService:%s' % self._path


class FiltersService(Service):
    """
    """

    def __init__(self, connection, path):
        super(FiltersService, self).__init__(connection, path)
        self._filter_service = None

    def add(
        self,
        filter,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, types.Filter),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(filter, headers, query, wait)

    def list(
        self,
        filter=None,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of filters to return. If not specified all the filters are returned.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def filter_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return FilterService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.filter_service(path)
        return self.filter_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'FiltersService:%s' % self._path


class GlusterBricksService(Service):
    """
    This service manages the gluster bricks in a gluster volume

    """

    def __init__(self, connection, path):
        super(GlusterBricksService, self).__init__(connection, path)
        self._brick_service = None

    def activate(
        self,
        async=None,
        bricks=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Activate the bricks post data migration of remove brick operation.
        Used to activate brick(s) once the data migration from bricks is complete but user no longer wishes to remove
        bricks. The bricks that were previously marked for removal will now be used as normal bricks.
        For example, to retain the bricks that on glustervolume `123` from which data was migrated, send a request like
        this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks/activate
        ----
        With a request body like this:
        [source,xml]
        ----
        <action>
          <bricks>
            <brick>
              <name>host1:/rhgs/brick1</name>
            </brick>
          </bricks>
        </action>
        ----


        This method supports the following parameters:

        `bricks`:: The list of bricks that need to be re-activated.

        `async`:: Indicates if the activation should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('bricks', bricks, list),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            bricks=bricks,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'activate', None, headers, query, wait)

    def add(
        self,
        bricks,
        replica_count=None,
        stripe_count=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adds a list of bricks to gluster volume.
        Used to expand a gluster volume by adding bricks. For replicated volume types, the parameter `replica_count`
        needs to be passed. In case the replica count is being increased, then the number of bricks needs to be
        equivalent to the number of replica sets.
        For example, to add bricks to gluster volume `123`, send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks
        ----
        With a request body like this:
        [source,xml]
        ----
        <bricks>
          <brick>
            <server_id>111</server_id>
            <brick_dir>/export/data/brick3</brick_dir>
          </brick>
        </bricks>
        ----


        This method supports the following parameters:

        `bricks`:: The list of bricks to be added to the volume

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('bricks', bricks, list),
            ('replica_count', replica_count, int),
            ('stripe_count', stripe_count, int),
        ])

        # Build the URL:
        query = query or {}
        if replica_count is not None:
            replica_count = Writer.render_integer(replica_count)
            query['replica_count'] = replica_count
        if stripe_count is not None:
            stripe_count = Writer.render_integer(stripe_count)
            query['stripe_count'] = stripe_count

        # Send the request and wait for the response:
        return self._internal_add(bricks, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists the bricks of a gluster volume.
        For example, to list bricks of gluster volume `123`, send a request like this:
        [source]
        ----
        GET /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks
        ----
        Provides an output as below:
        [source,xml]
        ----
        <bricks>
          <brick id="234">
            <name>host1:/rhgs/data/brick1</name>
            <brick_dir>/rhgs/data/brick1</brick_dir>
            <server_id>111</server_id>
            <status>up</status>
          </brick>
          <brick id="233">
            <name>host2:/rhgs/data/brick1</name>
            <brick_dir>/rhgs/data/brick1</brick_dir>
            <server_id>222</server_id>
            <status>up</status>
          </brick>
        </bricks>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of bricks to return. If not specified all the bricks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def migrate(
        self,
        async=None,
        bricks=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Start migration of data prior to removing bricks.
        Removing bricks is a two-step process, where the data on bricks to be removed, is first migrated to remaining
        bricks. Once migration is completed the removal of bricks is confirmed via the API
        <<services/gluster_bricks/methods/remove, remove>>. If at any point, the action needs to be cancelled
        <<services/gluster_bricks/methods/stop_migrate, stopmigrate>> has to be called.
        For instance, to delete a brick from a gluster volume with id `123`, send a request:
        [source]
        ----
        POST /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks/migrate
        ----
        With a request body like this:
        [source,xml]
        ----
        <action>
          <bricks>
            <brick>
              <name>host1:/rhgs/brick1</name>
            </brick>
          </bricks>
        </action>
        ----
        The migration process can be tracked from the job id returned from the API using
        <<services/job/methods/get, job>> and steps in job using <<services/step/methods/get, step>>


        This method supports the following parameters:

        `bricks`:: List of bricks for which data migration needs to be started.

        `async`:: Indicates if the migration should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('bricks', bricks, list),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            bricks=bricks,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'migrate', None, headers, query, wait)

    def remove(
        self,
        bricks=None,
        replica_count=None,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes bricks from gluster volume.
        The recommended way to remove bricks without data loss is to first migrate the data using
        <<services/gluster_bricks/methods/stop_migrate, stopmigrate>> and then removing them. If migrate was not called on
        bricks prior to remove, the bricks are removed without data migration which may lead to data loss.
        For example, to delete the bricks from gluster volume `123`, send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks
        ----
        With a request body like this:
        [source,xml]
        ----
        <bricks>
          <brick>
            <name>host:brick_directory</name>
          </brick>
        </bricks>
        ----


        This method supports the following parameters:

        `bricks`:: The list of bricks to be removed

        `replica_count`:: Replica count of volume post add operation.

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('bricks', bricks, list),
            ('replica_count', replica_count, int),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if bricks is not None:
            query['bricks'] = bricks
        if replica_count is not None:
            replica_count = Writer.render_integer(replica_count)
            query['replica_count'] = replica_count
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def stop_migrate(
        self,
        async=None,
        bricks=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Stops migration of data from bricks for a remove brick operation.
        To cancel data migration that was started as part of the 2-step remove brick process in case the user wishes to
        continue using the bricks. The bricks that were marked for removal will function as normal bricks post this
        operation.
        For example, to stop migration of data from the bricks of gluster volume `123`, send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks/stopmigrate
        ----
        With a request body like this:
        [source,xml]
        ----
        <bricks>
          <brick>
            <name>host:brick_directory</name>
          </brick>
        </bricks>
        ----


        This method supports the following parameters:

        `bricks`:: List of bricks for which data migration needs to be stopped. This list should match the arguments passed to
        <<services/gluster_bricks/methods/migrate, migrate>>.

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('bricks', bricks, list),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            bricks=bricks,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'stopmigrate', None, headers, query, wait)

    def brick_service(self, id):
        """
        Returns a reference to the service managing a single gluster brick.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return GlusterBrickService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.brick_service(path)
        return self.brick_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'GlusterBricksService:%s' % self._path


class GlusterHookService(Service):
    """
    """

    def __init__(self, connection, path):
        super(GlusterHookService, self).__init__(connection, path)

    def disable(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Resolves status conflict of hook among servers in cluster by disabling Gluster hook in all servers of the
        cluster. This updates the hook status to `DISABLED` in database.


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'disable', None, headers, query, wait)

    def enable(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Resolves status conflict of hook among servers in cluster by disabling Gluster hook in all servers of the
        cluster. This updates the hook status to `DISABLED` in database.


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'enable', None, headers, query, wait)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the this Gluster hook from all servers in cluster and deletes it from the database.


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def resolve(
        self,
        async=None,
        host=None,
        resolution_type=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Resolves missing hook conflict depending on the resolution type.
        For `ADD` resolves by copying hook stored in engine database to all servers where the hook is missing. The
        engine maintains a list of all servers where hook is missing.
        For `COPY` resolves conflict in hook content by copying hook stored in engine database to all servers where
        the hook is missing. The engine maintains a list of all servers where the content is conflicting. If a host
        id is passed as parameter, the hook content from the server is used as the master to copy to other servers
        in cluster.


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('host', host, types.Host),
            ('resolution_type', resolution_type, str),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            host=host,
            resolution_type=resolution_type,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'resolve', None, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'GlusterHookService:%s' % self._path


class GlusterHooksService(Service):
    """
    """

    def __init__(self, connection, path):
        super(GlusterHooksService, self).__init__(connection, path)
        self._hook_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of hooks to return. If not specified all the hooks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def hook_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return GlusterHookService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.hook_service(path)
        return self.hook_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'GlusterHooksService:%s' % self._path


class GlusterVolumesService(Service):
    """
    This service manages a collection of gluster volumes available in a cluster.

    """

    def __init__(self, connection, path):
        super(GlusterVolumesService, self).__init__(connection, path)
        self._volume_service = None

    def add(
        self,
        volume,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new gluster volume.
        The volume is created based on properties of the `volume` parameter. The properties `name`, `volume_type` and
        `bricks` are required.
        For example, to add a volume with name `myvolume` to the cluster `123`, send the following request:
        [source]
        ----
        POST /ovirt-engine/api/clusters/123/glustervolumes
        ----
        With the following request body:
        [source,xml]
        ----
        <gluster_volume>
          <name>myvolume</name>
          <volume_type>replicate</volume_type>
          <replica_count>3</replica_count>
          <bricks>
            <brick>
              <server_id>server1</server_id>
              <brick_dir>/exp1</brick_dir>
            </brick>
            <brick>
              <server_id>server2</server_id>
              <brick_dir>/exp1</brick_dir>
            </brick>
            <brick>
              <server_id>server3</server_id>
              <brick_dir>/exp1</brick_dir>
            </brick>
          <bricks>
        </gluster_volume>
        ----


        This method supports the following parameters:

        `volume`:: The gluster volume definition from which to create the volume is passed as input and the newly created
        volume is returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('volume', volume, types.GlusterVolume),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(volume, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists all gluster volumes in the cluster.
        For example, to list all Gluster Volumes in cluster `456`, send a request like
        this:
        [source]
        ----
        GET /ovirt-engine/api/clusters/456/glustervolumes
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of volumes to return. If not specified all the volumes are returned.

        `search`:: A query string used to restrict the returned volumes.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def volume_service(self, id):
        """
        Reference to a service managing gluster volume.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return GlusterVolumeService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.volume_service(path)
        return self.volume_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'GlusterVolumesService:%s' % self._path


class GroupService(Service):
    """
    """

    def __init__(self, connection, path):
        super(GroupService, self).__init__(connection, path)
        self._permissions_service = None
        self._roles_service = None
        self._tags_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def permissions_service(self):
        """
        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def roles_service(self):
        """
        """
        return AssignedRolesService(self._connection, '%s/roles' % self._path)

    def tags_service(self):
        """
        """
        return AssignedTagsService(self._connection, '%s/tags' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'roles':
            return self.roles_service()
        if path.startswith('roles/'):
            return self.roles_service().service(path[6:])
        if path == 'tags':
            return self.tags_service()
        if path.startswith('tags/'):
            return self.tags_service().service(path[5:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'GroupService:%s' % self._path


class GroupsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(GroupsService, self).__init__(connection, path)
        self._group_service = None

    def add(
        self,
        group,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add group from a directory service. Please note that domain name is name of the authorization provider.
        For example, to add the `Developers` group from the `internal-authz` authorization provider send a request
        like this:
        [source]
        ----
        POST /ovirt-engine/api/groups
        ----
        With a request body like this:
        [source,xml]
        ----
        <group>
          <name>Developers</name>
          <domain>
            <name>internal-authz</name>
          </domain>
        </group>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('group', group, types.Group),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(group, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of groups to return. If not specified all the groups are returned.

        `search`:: A query string used to restrict the returned groups.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def group_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return GroupService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.group_service(path)
        return self.group_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'GroupsService:%s' % self._path


class HostDeviceService(Service):
    """
    A service to access a particular device of a host.

    """

    def __init__(self, connection, path):
        super(HostDeviceService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieve information about a particular host's device.
        An example of getting a host device:
        [source]
        ----
        GET /ovirt-engine/api/hosts/123/devices/456
        ----
        [source,xml]
        ----
        <host_device href="/ovirt-engine/api/hosts/123/devices/456" id="456">
          <name>usb_1_9_1_1_0</name>
          <capability>usb</capability>
          <host href="/ovirt-engine/api/hosts/123" id="123"/>
          <parent_device href="/ovirt-engine/api/hosts/123/devices/789" id="789">
            <name>usb_1_9_1</name>
          </parent_device>
        </host_device>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'HostDeviceService:%s' % self._path


class HostDevicesService(Service):
    """
    A service to access host devices.

    """

    def __init__(self, connection, path):
        super(HostDevicesService, self).__init__(connection, path)
        self._device_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List the devices of a host.


        This method supports the following parameters:

        `max`:: Sets the maximum number of devices to return. If not specified all the devices are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def device_service(self, id):
        """
        Reference to the service that can be used to access a specific host device.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return HostDeviceService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.device_service(path)
        return self.device_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'HostDevicesService:%s' % self._path


class HostHookService(Service):
    """
    """

    def __init__(self, connection, path):
        super(HostHookService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'HostHookService:%s' % self._path


class HostHooksService(Service):
    """
    """

    def __init__(self, connection, path):
        super(HostHooksService, self).__init__(connection, path)
        self._hook_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of hooks to return. If not specified all the hooks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def hook_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return HostHookService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.hook_service(path)
        return self.hook_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'HostHooksService:%s' % self._path


class HostNicsService(Service):
    """
    A service to manage the network interfaces of a host.

    """

    def __init__(self, connection, path):
        super(HostNicsService, self).__init__(connection, path)
        self._nic_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of NICs to return. If not specified all the NICs are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def nic_service(self, id):
        """
        Reference to the service that manages a single network interface.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return HostNicService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.nic_service(path)
        return self.nic_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'HostNicsService:%s' % self._path


class HostNumaNodesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(HostNumaNodesService, self).__init__(connection, path)
        self._node_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of nodes to return. If not specified all the nodes are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def node_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return HostNumaNodeService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.node_service(path)
        return self.node_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'HostNumaNodesService:%s' % self._path


class HostStorageService(Service):
    """
    A service to manage host storages.

    """

    def __init__(self, connection, path):
        super(HostStorageService, self).__init__(connection, path)
        self._storage_service = None

    def list(
        self,
        report_status=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get list of storages.
        [source]
        ----
        GET /ovirt-engine/api/hosts/123/storage
        ----
        The XML response you get will be like this one:
        [source,xml]
        ----
        <host_storages>
          <host_storage id="123">
            ...
          </host_storage>
          ...
        </host_storages>
        ----


        This method supports the following parameters:

        `report_status`:: Indicates if the status of the LUNs in the storage should be checked.
        Checking the status of the LUN is an heavy weight operation and
        this data is not always needed by the user.
        This parameter will give the option to not perform the status check of the LUNs.
        The default is `true` for backward compatibility.
        Here an example with the LUN status :
        [source,xml]
        ----
        <host_storage id="123">
          <logical_units>
            <logical_unit id="123">
              <lun_mapping>0</lun_mapping>
              <paths>1</paths>
              <product_id>lun0</product_id>
              <serial>123</serial>
              <size>10737418240</size>
              <status>used</status>
              <vendor_id>LIO-ORG</vendor_id>
              <volume_group_id>123</volume_group_id>
            </logical_unit>
          </logical_units>
          <type>iscsi</type>
          <host id="123"/>
        </host_storage>
        ----
        Here an example without the LUN status :
        [source,xml]
        ----
        <host_storage id="123">
          <logical_units>
            <logical_unit id="123">
              <lun_mapping>0</lun_mapping>
              <paths>1</paths>
              <product_id>lun0</product_id>
              <serial>123</serial>
              <size>10737418240</size>
              <vendor_id>LIO-ORG</vendor_id>
              <volume_group_id>123</volume_group_id>
            </logical_unit>
          </logical_units>
          <type>iscsi</type>
          <host id="123"/>
        </host_storage>
        ----

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('report_status', report_status, bool),
        ])

        # Build the URL:
        query = query or {}
        if report_status is not None:
            report_status = Writer.render_boolean(report_status)
            query['report_status'] = report_status

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def storage_service(self, id):
        """
        Reference to a service managing the storage.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return StorageService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.storage_service(path)
        return self.storage_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'HostStorageService:%s' % self._path


class HostsService(Service):
    """
    A service that manages hosts.

    """

    def __init__(self, connection, path):
        super(HostsService, self).__init__(connection, path)
        self._host_service = None

    def add(
        self,
        host,
        deploy_hosted_engine=None,
        undeploy_hosted_engine=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new host.
        The host is created based on the attributes of the `host` parameter. The `name`, `address` and `root_password`
        properties are required.
        For example, to add a host send the following request:
        [source]
        ----
        POST /ovirt-engine/api/hosts
        ----
        With the following request body:
        [source,xml]
        ----
        <host>
          <name>myhost</name>
          <address>myhost.example.com</address>
          <root_password>myrootpassword</root_password>
        </host>
        ----
        NOTE: The `root_password` element is only included in the client-provided initial representation and is not
        exposed in the representations returned from subsequent requests.
        To add a hosted engine host, use the optional `deploy_hosted_engine` parameter:
        [source]
        ----
        POST /ovirt-engine/api/hosts?deploy_hosted_engine=true
        ----


        This method supports the following parameters:

        `host`:: The host definition from which to create the new host is passed as parameter, and the newly created host
        is returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('host', host, types.Host),
            ('deploy_hosted_engine', deploy_hosted_engine, bool),
            ('undeploy_hosted_engine', undeploy_hosted_engine, bool),
        ])

        # Build the URL:
        query = query or {}
        if deploy_hosted_engine is not None:
            deploy_hosted_engine = Writer.render_boolean(deploy_hosted_engine)
            query['deploy_hosted_engine'] = deploy_hosted_engine
        if undeploy_hosted_engine is not None:
            undeploy_hosted_engine = Writer.render_boolean(undeploy_hosted_engine)
            query['undeploy_hosted_engine'] = undeploy_hosted_engine

        # Send the request and wait for the response:
        return self._internal_add(host, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        filter=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get a list of all available hosts.
        For example, to list the hosts send the following request:
        ....
        GET /ovirt-engine/api/hosts
        ....
        The response body will be something like this:
        [source,xml]
        ----
        <hosts>
          <host href="/ovirt-engine/api/hosts/123" id="123">
            ...
          </host>
          <host href="/ovirt-engine/api/hosts/456" id="456">
            ...
          </host>
          ...
        </host>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of hosts to return. If not specified all the hosts are returned.

        `search`:: A query string used to restrict the returned hosts.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('filter', filter, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def host_service(self, id):
        """
        A Reference to service managing a specific host.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return HostService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.host_service(path)
        return self.host_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'HostsService:%s' % self._path


class IconService(Service):
    """
    A service to manage an icon (read-only).

    """

    def __init__(self, connection, path):
        super(IconService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get an icon.
        [source]
        ----
        GET /ovirt-engine/api/icons/123
        ----
        You will get a XML response like this one:
        [source,xml]
        ----
        <icon id="123">
          <data>Some binary data here</data>
          <media_type>image/png</media_type>
        </icon>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'IconService:%s' % self._path


class IconsService(Service):
    """
    A service to manage icons.

    """

    def __init__(self, connection, path):
        super(IconsService, self).__init__(connection, path)
        self._icon_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get a list of icons.
        [source]
        ----
        GET /ovirt-engine/api/icons
        ----
        You will get a XML response which is similar to this one:
        [source,xml]
        ----
        <icons>
          <icon id="123">
            <data>...</data>
            <media_type>image/png</media_type>
          </icon>
          ...
        </icons>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of icons to return. If not specified all the icons are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def icon_service(self, id):
        """
        Reference to the service that manages an specific icon.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return IconService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.icon_service(path)
        return self.icon_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'IconsService:%s' % self._path


class ImageService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ImageService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def import_(
        self,
        async=None,
        cluster=None,
        disk=None,
        import_as_template=None,
        storage_domain=None,
        template=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `cluster`:: Cluster where the image should be imported. Has effect only in case `import_as_template` parameter
        is set to `true`.

        `disk`:: The disk which should be imported.

        `import_as_template`:: Specify if template should be created from the imported disk.

        `template`:: Name of the template, which should be created. Has effect only in case `import_as_template` parameter
        is set to `true`.

        `storage_domain`:: Storage domain where disk should be imported.

        `async`:: Indicates if the import should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('cluster', cluster, types.Cluster),
            ('disk', disk, types.Disk),
            ('import_as_template', import_as_template, bool),
            ('storage_domain', storage_domain, types.StorageDomain),
            ('template', template, types.Template),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            cluster=cluster,
            disk=disk,
            import_as_template=import_as_template,
            storage_domain=storage_domain,
            template=template,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'import', None, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ImageService:%s' % self._path


class ImageTransferService(Service):
    """
    This service provides a mechanism to control an image transfer. The client will have
    to create a transfer by using <<services/image_transfers/methods/add, add>>
    of the <<services/image_transfers>> service, stating the image to transfer
    data to/from.
    After doing that, the transfer is managed by this service.
    E.g., for uploading to the disk image with id `52cb593f-837c-4633-a444-35a0a0383706`,
    the client can use oVirt's Python's SDK as follows:
    [source,python]
    ----
    transfers_service = system_service.image_transfers_service()
    transfer = transfers_service.add(
       types.ImageTransfer(
          image=types.Image(
             id='52cb593f-837c-4633-a444-35a0a0383706'
          )
       )
    )
    ----
    If the user wishes to download a disk rather than upload, he/she should specify
    `download` as the <<types/image_transfer_direction, direction>> attribute of the transfer.
    This will grant a read permission from the image, instead of a write permission.
    E.g:
    [source,python]
    ----
    transfers_service = system_service.image_transfers_service()
    transfer = transfers_service.add(
       types.ImageTransfer(
          image=types.Image(
             id='52cb593f-837c-4633-a444-35a0a0383706'
          ),
          direction=types.ImageTransferDirection.DOWNLOAD
       )
    )
    ----
    Transfers have phases, which govern the flow of the upload/download.
    A client implementing such a flow should poll/check the transfer's phase and
    act accordingly. All the possible phases can be found in
    <<types/image_transfer_phase, ImageTransferPhase>>.
    After adding a new transfer, its phase will be <<types/image_transfer_phase, initializing>>.
    The client will have to poll on the transfer's phase until it changes.
    When the phase becomes <<types/image_transfer_phase, transferring>>,
    the session is ready to start the transfer.
    For example:
    [source,python]
    ----
    transfer_service = transfers_service.image_transfer_service(transfer.id)
    while transfer.phase == types.ImageTransferPhase.INITIALIZING:
       time.sleep(3)
       transfer = transfer_service.get()
    ----
    At that stage, if the transfer's phase is <<types/image_transfer_phase, paused_system>>, then the session was
    not successfully established. One possible reason for that is that the ovirt-imageio-daemon is not running
    in the host that was selected for transfer.
    The transfer can be resumed by calling <<services/image_transfer/methods/resume, resume>>
    of the service that manages it.
    If the session was successfully established - the returned transfer entity will
    contain the <<types/image_transfer, proxy_url>> and <<types/image_transfer, signed_ticket>> attributes,
    which the client needs to use in order to transfer the required data. The client can choose whatever
    technique and tool for sending the HTTPS request with the image's data.
    - `proxy_url` is the address of a proxy server to the image, to do I/O to.
    - `signed_ticket` is the content that needs to be added to the `Authentication`
       header in the HTTPS request, in order to perform a trusted communication.
    For example, Python's HTTPSConnection can be used in order to perform a transfer,
    so an `transfer_headers` dict is set for the upcoming transfer:
    [source,python]
    ----
    transfer_headers = {
       'Authorization' :  transfer.signed_ticket,
    }
    ----
    Using Python's `HTTPSConnection`, a new connection is established:
    [source,python]
    ----
    # Extract the URI, port, and path from the transfer's proxy_url.
    url = urlparse.urlparse(transfer.proxy_url)
    # Create a new instance of the connection.
    proxy_connection = HTTPSConnection(
       url.hostname,
       url.port,
       context=ssl.SSLContext(ssl.PROTOCOL_SSLv23)
    )
    ----
    For upload, the specific content range being sent must be noted in the `Content-Range` HTTPS
    header. This can be used in order to split the transfer into several requests for
    a more flexible process.
    For doing that, the client will have to repeatedly extend the transfer session
    to keep the channel open. Otherwise, the session will terminate and the transfer will
    get into `paused_system` phase, and HTTPS requests to the server will be rejected.
    E.g., the client can iterate on chunks of the file, and send them to the
    proxy server while asking the service to extend the session:
    [source,python]
    ----
    path = "/path/to/image"
    MB_per_request = 32
    with open(path, "rb") as disk:
       size = os.path.getsize(path)
       chunk_size = 1024*1024*MB_per_request
       pos = 0
       while (pos < size):
          transfer_service.extend()
          transfer_headers['Content-Range'] = "bytes %d-%d/%d" % (pos, min(pos + chunk_size, size)-1, size)
          proxy_connection.request(
             'PUT',
             url.path,
             disk.read(chunk_size),
             headers=transfer_headers
          )
          r = proxy_connection.getresponse()
          print r.status, r.reason, "Completed", "{:.0%}".format(pos/ float(size))
          pos += chunk_size
    ----
    Similarly, for a download transfer, a `Range` header must be sent, making the download process
    more easily managed by downloading the disk in chunks.
    E.g., the client will again iterate on chunks of the disk image, but this time he/she will download
    it to a local file, rather than uploading its own file to the image:
    [source,python]
    ----
    output_file = "/home/user/downloaded_image"
    MiB_per_request = 32
    chunk_size = 1024*1024*MiB_per_request
    total = disk_size
    with open(output_file, "wb") as disk:
       pos = 0
       while pos < total:
          transfer_service.extend()
          transfer_headers['Range'] = "bytes=%d-%d" %  (pos, min(total, pos + chunk_size) - 1)
          proxy_connection.request('GET', proxy_url.path, headers=transfer_headers)
          r = proxy_connection.getresponse()
          disk.write(r.read())
          print "Completed", "{:.0%}".format(pos/ float(total))
          pos += chunk_size
    ----
    When finishing the transfer, the user should call
    <<services/image_transfer/methods/finalize, finalize>>. This will make the
    final adjustments and verifications for finishing the transfer process.
    For example:
    [source,python]
    ----
    transfer_service.finalize()
    ----
    In case of an error, the transfer's phase will be changed to
    <<types/image_transfer_phase, finished_failure>>, and
    the disk's status will be changed to `Illegal`. Otherwise it will be changed to
    <<types/image_transfer_phase, finished_success>>, and the disk will be ready
    to be used. In both cases, the transfer entity will be removed shortly after.

    """

    def __init__(self, connection, path):
        super(ImageTransferService, self).__init__(connection, path)

    def extend(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Extend the image transfer session.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'extend', None, headers, query, wait)

    def finalize(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        After finishing to transfer the data, finalize the transfer.
        This will make sure that the data being transferred is valid and fits the
        image entity that was targeted in the transfer. Specifically, will verify that
        if the image entity is a QCOW disk, the data uploaded is indeed a QCOW file,
        and that the image doesn't have a backing file.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'finalize', None, headers, query, wait)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get the image transfer entity.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def pause(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Pause the image transfer session.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'pause', None, headers, query, wait)

    def resume(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Resume the image transfer session. The client will need to poll the transfer's phase until
        it is different than `resuming`. For example:
        [source,python]
        ----
        transfer_service = transfers_service.image_transfer_service(transfer.id)
        transfer_service.resume()
        transfer = transfer_service.get()
        while transfer.phase == types.ImageTransferPhase.RESUMING:
           time.sleep(1)
           transfer = transfer_service.get()
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'resume', None, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ImageTransferService:%s' % self._path


class ImageTransfersService(Service):
    """
    This service manages image transfers, for performing Image I/O API in oVirt.
    Please refer to <<services/image_transfer, image transfer>> for further
    documentation.

    """

    def __init__(self, connection, path):
        super(ImageTransfersService, self).__init__(connection, path)
        self._image_transfer_service = None

    def add(
        self,
        image_transfer,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add a new image transfer. An image needs to be specified in order to make
        a new transfer.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('image_transfer', image_transfer, types.ImageTransfer),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(image_transfer, headers, query, wait)

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the list of image transfers that are currently
        being performed.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def image_transfer_service(self, id):
        """
        Returns a reference to the service that manages an
        specific image transfer.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return ImageTransferService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.image_transfer_service(path)
        return self.image_transfer_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'ImageTransfersService:%s' % self._path


class ImagesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(ImagesService, self).__init__(connection, path)
        self._image_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of images to return. If not specified all the images are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def image_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return ImageService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.image_service(path)
        return self.image_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'ImagesService:%s' % self._path


class InstanceTypeService(Service):
    """
    """

    def __init__(self, connection, path):
        super(InstanceTypeService, self).__init__(connection, path)
        self._graphics_consoles_service = None
        self._nics_service = None
        self._watchdogs_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get a specific instance type and it's attributes.
        [source]
        ----
        GET /ovirt-engine/api/instancetypes/123
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a specific instance type from the system.
        If a virtual machine was created using an instance type X after removal of the instance type
        the virtual machine's instance type will be set to `custom`.
        [source]
        ----
        DELETE /ovirt-engine/api/instancetypes/123
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        instance_type,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Update a specific instance type and it's attributes.
        All the attributes are editable after creation.
        If a virtual machine was created using an instance type X and some configuration in instance
        type X was updated, the virtual machine's configuration will be updated automatically by the
        engine.
        [source]
        ----
        PUT /ovirt-engine/api/instancetypes/123
        ----
        For example, to update the memory of instance type `123` to 1 GiB and set the cpu topology
        to 2 sockets and 1 core, send a request like this:
        [source, xml]
        ----
        <instance_type>
          <memory>1073741824</memory>
          <cpu>
            <topology>
              <cores>1</cores>
              <sockets>2</sockets>
              <threads>1</threads>
            </topology>
          </cpu>
        </instance_type>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('instance_type', instance_type, types.InstanceType),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(instance_type, headers, query, wait)

    def graphics_consoles_service(self):
        """
        Reference to the service that manages the graphic consoles that are attached to this
        instance type.

        """
        return InstanceTypeGraphicsConsolesService(self._connection, '%s/graphicsconsoles' % self._path)

    def nics_service(self):
        """
        Reference to the service that manages the NICs that are attached to this instance type.

        """
        return InstanceTypeNicsService(self._connection, '%s/nics' % self._path)

    def watchdogs_service(self):
        """
        Reference to the service that manages the watchdogs that are attached to this instance type.

        """
        return InstanceTypeWatchdogsService(self._connection, '%s/watchdogs' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'graphicsconsoles':
            return self.graphics_consoles_service()
        if path.startswith('graphicsconsoles/'):
            return self.graphics_consoles_service().service(path[17:])
        if path == 'nics':
            return self.nics_service()
        if path.startswith('nics/'):
            return self.nics_service().service(path[5:])
        if path == 'watchdogs':
            return self.watchdogs_service()
        if path.startswith('watchdogs/'):
            return self.watchdogs_service().service(path[10:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'InstanceTypeService:%s' % self._path


class InstanceTypeGraphicsConsoleService(Service):
    """
    """

    def __init__(self, connection, path):
        super(InstanceTypeGraphicsConsoleService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets graphics console configuration of the instance type.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove the graphics console from the instance type.


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'InstanceTypeGraphicsConsoleService:%s' % self._path


class InstanceTypeGraphicsConsolesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(InstanceTypeGraphicsConsolesService, self).__init__(connection, path)
        self._console_service = None

    def add(
        self,
        console,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add new graphics console to the instance type.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('console', console, types.GraphicsConsole),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(console, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists all the configured graphics consoles of the instance type.


        This method supports the following parameters:

        `max`:: Sets the maximum number of consoles to return. If not specified all the consoles are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def console_service(self, id):
        """
        Returns a reference to the service that manages a specific instance type graphics console.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return InstanceTypeGraphicsConsoleService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.console_service(path)
        return self.console_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'InstanceTypeGraphicsConsolesService:%s' % self._path


class InstanceTypeNicService(Service):
    """
    """

    def __init__(self, connection, path):
        super(InstanceTypeNicService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets network interface configuration of the instance type.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove the network interface from the instance type.


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        nic,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the network interface configuration of the instance type.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('nic', nic, types.Nic),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(nic, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'InstanceTypeNicService:%s' % self._path


class InstanceTypeNicsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(InstanceTypeNicsService, self).__init__(connection, path)
        self._nic_service = None

    def add(
        self,
        nic,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add new network interface to the instance type.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('nic', nic, types.Nic),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(nic, headers, query, wait)

    def list(
        self,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists all the configured network interface of the instance type.


        This method supports the following parameters:

        `max`:: Sets the maximum number of NICs to return. If not specified all the NICs are returned.

        `search`:: A query string used to restrict the returned templates.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def nic_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return InstanceTypeNicService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.nic_service(path)
        return self.nic_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'InstanceTypeNicsService:%s' % self._path


class InstanceTypeWatchdogService(Service):
    """
    """

    def __init__(self, connection, path):
        super(InstanceTypeWatchdogService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets watchdog configuration of the instance type.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove a watchdog from the instance type.


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        watchdog,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the watchdog configuration of the instance type.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('watchdog', watchdog, types.Watchdog),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(watchdog, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'InstanceTypeWatchdogService:%s' % self._path


class InstanceTypeWatchdogsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(InstanceTypeWatchdogsService, self).__init__(connection, path)
        self._watchdog_service = None

    def add(
        self,
        watchdog,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add new watchdog to the instance type.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('watchdog', watchdog, types.Watchdog),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(watchdog, headers, query, wait)

    def list(
        self,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists all the configured watchdogs of the instance type.


        This method supports the following parameters:

        `max`:: Sets the maximum number of watchdogs to return. If not specified all the watchdogs are
        returned.

        `search`:: A query string used to restrict the returned templates.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def watchdog_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return InstanceTypeWatchdogService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.watchdog_service(path)
        return self.watchdog_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'InstanceTypeWatchdogsService:%s' % self._path


class InstanceTypesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(InstanceTypesService, self).__init__(connection, path)
        self._instance_type_service = None

    def add(
        self,
        instance_type,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new instance type.
        This requires only a name attribute and can include all hardware configurations of the
        virtual machine.
        [source]
        ----
        POST /ovirt-engine/api/instancetypes
        ----
        With a request body like this:
        [source,xml]
        ----
        <instance_type>
          <name>myinstancetype</name>
        </template>
        ----
        Creating an instance type with all hardware configurations with a request body like this:
        [source,xml]
        ----
        <instance_type>
          <name>myinstancetype</name>
          <console>
            <enabled>true</enabled>
          </console>
          <cpu>
            <topology>
              <cores>2</cores>
              <sockets>2</sockets>
              <threads>1</threads>
            </topology>
          </cpu>
          <custom_cpu_model>AMD Opteron_G2</custom_cpu_model>
          <custom_emulated_machine>q35</custom_emulated_machine>
          <display>
            <monitors>1</monitors>
            <single_qxl_pci>true</single_qxl_pci>
            <smartcard_enabled>true</smartcard_enabled>
            <type>spice</type>
          </display>
          <high_availability>
            <enabled>true</enabled>
            <priority>1</priority>
          </high_availability>
          <io>
            <threads>2</threads>
          </io>
          <memory>4294967296</memory>
          <memory_policy>
            <ballooning>true</ballooning>
            <guaranteed>268435456</guaranteed>
          </memory_policy>
          <migration>
            <auto_converge>inherit</auto_converge>
            <compressed>inherit</compressed>
            <policy id="00000000-0000-0000-0000-000000000000"/>
          </migration>
          <migration_downtime>2</migration_downtime>
          <os>
            <boot>
              <devices>
                <device>hd</device>
              </devices>
            </boot>
          </os>
          <rng_device>
            <rate>
              <bytes>200</bytes>
              <period>2</period>
            </rate>
            <source>urandom</source>
          </rng_device>
          <soundcard_enabled>true</soundcard_enabled>
          <usb>
            <enabled>true</enabled>
            <type>native</type>
          </usb>
          <virtio_scsi>
            <enabled>true</enabled>
          </virtio_scsi>
        </instance_type>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('instance_type', instance_type, types.InstanceType),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(instance_type, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists all existing instance types in the system.


        This method supports the following parameters:

        `max`:: Sets the maximum number of instance types to return. If not specified all the instance
        types are returned.

        `search`:: A query string used to restrict the returned templates.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed
        taking case into account. The default value is `true`, which means that case is taken
        into account. If you want to search ignoring case set it to `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def instance_type_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return InstanceTypeService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.instance_type_service(path)
        return self.instance_type_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'InstanceTypesService:%s' % self._path


class IscsiBondService(Service):
    """
    """

    def __init__(self, connection, path):
        super(IscsiBondService, self).__init__(connection, path)
        self._networks_service = None
        self._storage_server_connections_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes of an existing iSCSI bond.
        For example, to remove the iSCSI bond `456` send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/datacenters/123/iscsibonds/456
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        bond,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates an iSCSI bond.
        Updating of an iSCSI bond can be done on the `name` and the `description` attributes only. For example, to
        update the iSCSI bond `456` of data center `123`, send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/datacenters/123/iscsibonds/1234
        ----
        The request body should look like this:
        [source,xml]
        ----
        <iscsi_bond>
           <name>mybond</name>
           <description>My iSCSI bond</description>
        </iscsi_bond>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('bond', bond, types.IscsiBond),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(bond, headers, query, wait)

    def networks_service(self):
        """
        """
        return NetworksService(self._connection, '%s/networks' % self._path)

    def storage_server_connections_service(self):
        """
        """
        return StorageServerConnectionsService(self._connection, '%s/storageserverconnections' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'networks':
            return self.networks_service()
        if path.startswith('networks/'):
            return self.networks_service().service(path[9:])
        if path == 'storageserverconnections':
            return self.storage_server_connections_service()
        if path.startswith('storageserverconnections/'):
            return self.storage_server_connections_service().service(path[25:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'IscsiBondService:%s' % self._path


class IscsiBondsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(IscsiBondsService, self).__init__(connection, path)
        self._iscsi_bond_service = None

    def add(
        self,
        bond,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Create a new iSCSI bond on a data center.
        For example, to create a new iSCSI bond on data center `123` using storage connections `456` and `789`, send a
        request like this:
        [source]
        ----
        POST /ovirt-engine/api/datacenters/123/iscsibonds
        ----
        The request body should look like this:
        [source,xml]
        ----
        <iscsi_bond>
          <name>mybond</name>
          <storage_connections>
            <storage_connection id="456"/>
            <storage_connection id="789"/>
          </storage_connections>
          <networks>
            <network id="abc"/>
          </networks>
        </iscsi_bond>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('bond', bond, types.IscsiBond),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(bond, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of bonds to return. If not specified all the bonds are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def iscsi_bond_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return IscsiBondService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.iscsi_bond_service(path)
        return self.iscsi_bond_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'IscsiBondsService:%s' % self._path


class JobService(Service):
    """
    A service to manage a job.

    """

    def __init__(self, connection, path):
        super(JobService, self).__init__(connection, path)
        self._steps_service = None

    def clear(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Set an external job execution to be cleared by the system.
        For example, to set a job with identifier `123` send the following request:
        [source]
        ----
        POST /ovirt-engine/api/jobs/clear
        ----
        With the following request body:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'clear', None, headers, query, wait)

    def end(
        self,
        async=None,
        force=None,
        succeeded=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Marks an external job execution as ended.
        For example, to terminate a job with identifier `123` send the following request:
        [source]
        ----
        POST /ovirt-engine/api/jobs/end
        ----
        With the following request body:
        [source,xml]
        ----
        <action>
          <force>true</force>
          <status>finished</status>
        </action>
        ----


        This method supports the following parameters:

        `force`:: Indicates if the job should be forcibly terminated.

        `succeeded`:: Indicates if the job should be marked as successfully finished or as failed.
        This parameter is optional, and the default value is `true`.

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('force', force, bool),
            ('succeeded', succeeded, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            force=force,
            succeeded=succeeded,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'end', None, headers, query, wait)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves a job.
        [source]
        ----
        GET /ovirt-engine/api/jobs/123
        ----
        You will receive response in XML like this one:
        [source,xml]
        ----
        <job href="/ovirt-engine/api/jobs/123" id="123">
          <actions>
            <link href="/ovirt-engine/api/jobs/123/clear" rel="clear"/>
            <link href="/ovirt-engine/api/jobs/123/end" rel="end"/>
          </actions>
          <description>Adding Disk</description>
          <link href="/ovirt-engine/api/jobs/123/steps" rel="steps"/>
          <auto_cleared>true</auto_cleared>
          <end_time>2016-12-12T23:07:29.758+02:00</end_time>
          <external>false</external>
          <last_updated>2016-12-12T23:07:29.758+02:00</last_updated>
          <start_time>2016-12-12T23:07:26.593+02:00</start_time>
          <status>failed</status>
          <owner href="/ovirt-engine/api/users/456" id="456"/>
        </job>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def steps_service(self):
        """
        List all the steps of the job.

        """
        return StepsService(self._connection, '%s/steps' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'steps':
            return self.steps_service()
        if path.startswith('steps/'):
            return self.steps_service().service(path[6:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'JobService:%s' % self._path


class JobsService(Service):
    """
    A service to manage jobs.

    """

    def __init__(self, connection, path):
        super(JobsService, self).__init__(connection, path)
        self._job_service = None

    def add(
        self,
        job,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add an external job.
        For example, to add a job with the following request:
        [source]
        ----
        POST /ovirt-engine/api/jobs
        ----
        With the following request body:
        [source,xml]
        ----
        <job>
          <description>Doing some work</description>
          <auto_cleared>true</auto_cleared>
        </job>
        ----
        The response should look like:
        [source,xml]
        ----
        <job href="/ovirt-engine/api/jobs/123" id="123">
          <actions>
            <link href="/ovirt-engine/api/jobs/123/clear" rel="clear"/>
            <link href="/ovirt-engine/api/jobs/123/end" rel="end"/>
          </actions>
          <description>Doing some work</description>
          <link href="/ovirt-engine/api/jobs/123/steps" rel="steps"/>
          <auto_cleared>true</auto_cleared>
          <external>true</external>
          <last_updated>2016-12-13T02:15:42.130+02:00</last_updated>
          <start_time>2016-12-13T02:15:42.130+02:00</start_time>
          <status>started</status>
          <owner href="/ovirt-engine/api/users/456" id="456"/>
        </job>
        ----


        This method supports the following parameters:

        `job`:: Job that will be added.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('job', job, types.Job),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(job, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the representation of the jobs.
        [source]
        ----
        GET /ovirt-engine/api/jobs
        ----
        You will receive response in XML like this one:
        [source,xml]
        ----
        <jobs>
          <job href="/ovirt-engine/api/jobs/123" id="123">
            <actions>
              <link href="/ovirt-engine/api/jobs/123/clear" rel="clear"/>
              <link href="/ovirt-engine/api/jobs/123/end" rel="end"/>
            </actions>
            <description>Adding Disk</description>
            <link href="/ovirt-engine/api/jobs/123/steps" rel="steps"/>
            <auto_cleared>true</auto_cleared>
            <end_time>2016-12-12T23:07:29.758+02:00</end_time>
            <external>false</external>
            <last_updated>2016-12-12T23:07:29.758+02:00</last_updated>
            <start_time>2016-12-12T23:07:26.593+02:00</start_time>
            <status>failed</status>
            <owner href="/ovirt-engine/api/users/456" id="456"/>
          </job>
          ...
        </jobs>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of jobs to return. If not specified all the jobs are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def job_service(self, id):
        """
        Reference to the job service.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return JobService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.job_service(path)
        return self.job_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'JobsService:%s' % self._path


class KatelloErrataService(Service):
    """
    A service to manage Katello errata.
    The information is retrieved from Katello.

    """

    def __init__(self, connection, path):
        super(KatelloErrataService, self).__init__(connection, path)
        self._katello_erratum_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the representation of the Katello errata.
        [source]
        ----
        GET /ovirt-engine/api/katelloerrata
        ----
        You will receive response in XML like this one:
        [source,xml]
        ----
        <katello_errata>
          <katello_erratum href="/ovirt-engine/api/katelloerrata/123" id="123">
            <name>RHBA-2013:XYZ</name>
            <description>The description of the erratum</description>
            <title>some bug fix update</title>
            <type>bugfix</type>
            <issued>2013-11-20T02:00:00.000+02:00</issued>
            <solution>Few guidelines regarding the solution</solution>
            <summary>Updated packages that fix one bug are now available for XYZ</summary>
            <packages>
              <package>
                <name>libipa_hbac-1.9.2-82.11.el6_4.i686</name>
              </package>
              ...
            </packages>
          </katello_erratum>
          ...
        </katello_errata>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of errata to return. If not specified all the errata are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def katello_erratum_service(self, id):
        """
        Reference to the Katello erratum service.
        Use this service to view the erratum by its id.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return KatelloErratumService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.katello_erratum_service(path)
        return self.katello_erratum_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'KatelloErrataService:%s' % self._path


class KatelloErratumService(Service):
    """
    A service to manage a Katello erratum.

    """

    def __init__(self, connection, path):
        super(KatelloErratumService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves a Katello erratum.
        [source]
        ----
        GET /ovirt-engine/api/katelloerrata/123
        ----
        You will receive response in XML like this one:
        [source,xml]
        ----
        <katello_erratum href="/ovirt-engine/api/katelloerrata/123" id="123">
          <name>RHBA-2013:XYZ</name>
          <description>The description of the erratum</description>
          <title>some bug fix update</title>
          <type>bugfix</type>
          <issued>2013-11-20T02:00:00.000+02:00</issued>
          <solution>Few guidelines regarding the solution</solution>
          <summary>Updated packages that fix one bug are now available for XYZ</summary>
          <packages>
            <package>
              <name>libipa_hbac-1.9.2-82.11.el6_4.i686</name>
            </package>
            ...
          </packages>
        </katello_erratum>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'KatelloErratumService:%s' % self._path


class MacPoolService(Service):
    """
    """

    def __init__(self, connection, path):
        super(MacPoolService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a MAC address pool.
        For example, to remove the MAC address pool having id `123` send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/macpools/123
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        pool,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates a MAC address pool.
        The `name`, `description`, `allow_duplicates`, and `ranges` attributes can be updated.
        For example, to update the MAC address pool of id `123` send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/macpools/123
        ----
        With a request body like this:
        [source,xml]
        ----
        <mac_pool>
          <name>UpdatedMACPool</name>
          <description>An updated MAC address pool</description>
          <allow_duplicates>false</allow_duplicates>
          <ranges>
            <range>
              <from>00:1A:4A:16:01:51</from>
              <to>00:1A:4A:16:01:e6</to>
            </range>
            <range>
              <from>02:1A:4A:01:00:00</from>
              <to>02:1A:4A:FF:FF:FF</to>
            </range>
          </ranges>
        </mac_pool>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('pool', pool, types.MacPool),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(pool, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'MacPoolService:%s' % self._path


class MacPoolsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(MacPoolsService, self).__init__(connection, path)
        self._mac_pool_service = None

    def add(
        self,
        pool,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new MAC address pool.
        Creation of a MAC address pool requires values for the `name` and `ranges` attributes.
        For example, to create MAC address pool send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/macpools
        ----
        With a request body like this:
        [source,xml]
        ----
        <mac_pool>
          <name>MACPool</name>
          <description>A MAC address pool</description>
          <allow_duplicates>true</allow_duplicates>
          <default_pool>false</default_pool>
          <ranges>
            <range>
              <from>00:1A:4A:16:01:51</from>
              <to>00:1A:4A:16:01:e6</to>
            </range>
          </ranges>
        </mac_pool>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('pool', pool, types.MacPool),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(pool, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of pools to return. If not specified all the pools are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def mac_pool_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return MacPoolService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.mac_pool_service(path)
        return self.mac_pool_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'MacPoolsService:%s' % self._path


class MeasurableService(Service):
    """
    """

    def __init__(self, connection, path):
        super(MeasurableService, self).__init__(connection, path)
        self._statistics_service = None

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'MeasurableService:%s' % self._path


class MoveableService(Service):
    """
    """

    def __init__(self, connection, path):
        super(MoveableService, self).__init__(connection, path)

    def move(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the move should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'move', None, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'MoveableService:%s' % self._path


class NetworkService(Service):
    """
    A service managing a network

    """

    def __init__(self, connection, path):
        super(NetworkService, self).__init__(connection, path)
        self._network_labels_service = None
        self._permissions_service = None
        self._vnic_profiles_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets a logical network.
        For example:
        [source]
        ----
        GET /ovirt-engine/api/networks/123
        ----
        Will respond:
        [source,xml]
        ----
        <network href="/ovirt-engine/api/networks/123" id="123">
          <name>ovirtmgmt</name>
          <description>Default Management Network</description>
          <link href="/ovirt-engine/api/networks/123/permissions" rel="permissions"/>
          <link href="/ovirt-engine/api/networks/123/vnicprofiles" rel="vnicprofiles"/>
          <link href="/ovirt-engine/api/networks/123/networklabels" rel="networklabels"/>
          <mtu>0</mtu>
          <stp>false</stp>
          <usages>
            <usage>vm</usage>
          </usages>
          <data_center href="/ovirt-engine/api/datacenters/456" id="456"/>
        </network>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a logical network, or the association of a logical network to a data center.
        For example, to remove the logical network `123` send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/networks/123
        ----
        Each network is bound exactly to one data center. So if we disassociate network with data center it has the same
        result as if we would just remove that network. However it might be more specific to say we're removing network
        `456` of data center `123`.
        For example, to remove the association of network `456` to data center `123` send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/datacenters/123/networks/456
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        network,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates a logical network.
        The `name`, `description`, `ip`, `vlan`, `stp` and `display` attributes can be updated.
        For example, to update the description of the logical network `123` send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/networks/123
        ----
        With a request body like this:
        [source,xml]
        ----
        <network>
          <description>My updated description</description>
        </network>
        ----
        The maximum transmission unit of a network is set using a PUT request to
        specify the integer value of the `mtu` attribute.
        For example, to set the maximum transmission unit send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/datacenters/123/networks/456
        ----
        With a request body like this:
        [source,xml]
        ----
        <network>
          <mtu>1500</mtu>
        </network>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('network', network, types.Network),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(network, headers, query, wait)

    def network_labels_service(self):
        """
        Reference to the service that manages the network labels assigned to this network.

        """
        return NetworkLabelsService(self._connection, '%s/networklabels' % self._path)

    def permissions_service(self):
        """
        Reference to the service that manages the permissions assigned to this network.

        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def vnic_profiles_service(self):
        """
        Reference to the service that manages the vNIC profiles assigned to this network.

        """
        return AssignedVnicProfilesService(self._connection, '%s/vnicprofiles' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'networklabels':
            return self.network_labels_service()
        if path.startswith('networklabels/'):
            return self.network_labels_service().service(path[14:])
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'vnicprofiles':
            return self.vnic_profiles_service()
        if path.startswith('vnicprofiles/'):
            return self.vnic_profiles_service().service(path[13:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'NetworkService:%s' % self._path


class NetworkAttachmentService(Service):
    """
    """

    def __init__(self, connection, path):
        super(NetworkAttachmentService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        attachment,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('attachment', attachment, types.NetworkAttachment),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(attachment, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'NetworkAttachmentService:%s' % self._path


class NetworkAttachmentsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(NetworkAttachmentsService, self).__init__(connection, path)
        self._attachment_service = None

    def add(
        self,
        attachment,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('attachment', attachment, types.NetworkAttachment),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(attachment, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of attachments to return. If not specified all the attachments are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def attachment_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return NetworkAttachmentService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.attachment_service(path)
        return self.attachment_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'NetworkAttachmentsService:%s' % self._path


class NetworkFilterService(Service):
    """
    Manages a network filter.
    [source,xml]
    ----
    <network_filter id="00000019-0019-0019-0019-00000000026b">
      <name>example-network-filter-b</name>
      <version>
        <major>4</major>
        <minor>0</minor>
        <build>-1</build>
        <revision>-1</revision>
      </version>
    </network_filter>
    ----
    Please note that version is referring to the minimal support version for the specific filter.

    """

    def __init__(self, connection, path):
        super(NetworkFilterService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves a representation of the network filter.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'NetworkFilterService:%s' % self._path


class NetworkFilterParameterService(Service):
    """
    This service manages a parameter for a network filter.

    """

    def __init__(self, connection, path):
        super(NetworkFilterParameterService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves a representation of the network filter parameter.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the filter parameter.
        For example, to remove the filter parameter with id `123` on NIC `456` of virtual machine `789`
        send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/vms/789/nics/456/networkfilterparameters/123
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        parameter,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the network filter parameter.
        For example, to update the network filter parameter having with with id `123` on NIC `456` of
        virtual machine `789` send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/vms/789/nics/456/networkfilterparameters/123
        ----
        With a request body like this:
        [source,xml]
        ----
        <network_filter_parameter>
          <name>updatedName</name>
          <value>updatedValue</value>
        </network_filter_parameter>
        ----


        This method supports the following parameters:

        `parameter`:: The network filter parameter that is being updated.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('parameter', parameter, types.NetworkFilterParameter),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_update(parameter, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'NetworkFilterParameterService:%s' % self._path


class NetworkFilterParametersService(Service):
    """
    This service manages a collection of parameters for network filters.

    """

    def __init__(self, connection, path):
        super(NetworkFilterParametersService, self).__init__(connection, path)
        self._parameter_service = None

    def add(
        self,
        parameter,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add a network filter parameter.
        For example, to add the parameter for the network filter on NIC `456` of
        virtual machine `789` send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/vms/789/nics/456/networkfilterparameters
        ----
        With a request body like this:
        [source,xml]
        ----
        <network_filter_parameter>
          <name>IP</name>
          <value>10.0.1.2</value>
        </network_filter_parameter>
        ----


        This method supports the following parameters:

        `parameter`:: The network filter parameter that is being added.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('parameter', parameter, types.NetworkFilterParameter),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(parameter, headers, query, wait)

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the representations of the network filter parameters.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def parameter_service(self, id):
        """
        Reference to the service that manages a specific network filter parameter.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return NetworkFilterParameterService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.parameter_service(path)
        return self.parameter_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'NetworkFilterParametersService:%s' % self._path


class NetworkFiltersService(Service):
    """
    Represents a readonly network filters sub-collection.
    The network filter enables to filter packets send to/from the VM's nic according to defined rules.
    For more information please refer to <<services/network_filter,NetworkFilter>> service documentation
    Network filters are supported in different versions, starting from version 3.0.
    A network filter is defined for each vnic profile.
    A vnic profile is defined for a specific network.
    A network can be assigned to several different clusters. In the future, each network will be defined in
    cluster level.
    Currently, each network is being defined at data center level. Potential network filters for each network
    are determined by the network's data center compatibility version V.
    V must be >= the network filter version in order to configure this network filter for a specific network.
    Please note, that if a network is assigned to cluster with a version supporting a network filter, the filter
    may not be available due to the data center version being smaller then the network filter's version.
    Example of listing all of the supported network filters for a specific cluster:
    [source]
    ----
    GET http://localhost:8080/ovirt-engine/api/clusters/{cluster:id}/networkfilters
    ----
    Output:
    [source,xml]
    ----
    <network_filters>
      <network_filter id="00000019-0019-0019-0019-00000000026c">
        <name>example-network-filter-a</name>
        <version>
          <major>4</major>
          <minor>0</minor>
          <build>-1</build>
          <revision>-1</revision>
        </version>
      </network_filter>
      <network_filter id="00000019-0019-0019-0019-00000000026b">
        <name>example-network-filter-b</name>
        <version>
          <major>4</major>
          <minor>0</minor>
          <build>-1</build>
          <revision>-1</revision>
        </version>
      </network_filter>
      <network_filter id="00000019-0019-0019-0019-00000000026a">
        <name>example-network-filter-a</name>
        <version>
          <major>3</major>
          <minor>0</minor>
          <build>-1</build>
          <revision>-1</revision>
        </version>
      </network_filter>
    </network_filters>
    ----

    """

    def __init__(self, connection, path):
        super(NetworkFiltersService, self).__init__(connection, path)
        self._network_filter_service = None

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the representations of the network filters.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def network_filter_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return NetworkFilterService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.network_filter_service(path)
        return self.network_filter_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'NetworkFiltersService:%s' % self._path


class NetworkLabelService(Service):
    """
    """

    def __init__(self, connection, path):
        super(NetworkLabelService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a label from a logical network.
        For example, to remove the label `exemplary` from a logical network having id `123` send the following request:
        [source]
        ----
        DELETE /ovirt-engine/api/networks/123/labels/exemplary
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'NetworkLabelService:%s' % self._path


class NetworkLabelsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(NetworkLabelsService, self).__init__(connection, path)
        self._label_service = None

    def add(
        self,
        label,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Attaches label to logical network.
        You can attach labels to a logical network to automate the association of that logical network with physical host
        network interfaces to which the same label has been attached.
        For example, to attach the label `mylabel` to a logical network having id `123` send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/networks/123/labels
        ----
        With a request body like this:
        [source,xml]
        ----
        <label id="mylabel"/>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('label', label, types.NetworkLabel),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(label, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of labels to return. If not specified all the labels are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def label_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return NetworkLabelService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.label_service(path)
        return self.label_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'NetworkLabelsService:%s' % self._path


class NetworksService(Service):
    """
    Manages logical networks.
    The engine creates a default `ovirtmgmt` network on installation. This network acts as the management network for
    access to hypervisor hosts. This network is associated with the `Default` cluster and is a member of the `Default`
    data center.

    """

    def __init__(self, connection, path):
        super(NetworksService, self).__init__(connection, path)
        self._network_service = None

    def add(
        self,
        network,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new logical network, or associates an existing network with a data center.
        Creation of a new network requires the `name` and `data_center` elements.
        For example, to create a network named `mynetwork` for data center `123` send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/networks
        ----
        With a request body like this:
        [source,xml]
        ----
        <network>
          <name>mynetwork</name>
          <data_center id="123"/>
        </network>
        ----
        To associate the existing network `456` with the data center `123` send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/datacenters/123/networks
        ----
        With a request body like this:
        [source,xml]
        ----
        <network>
          <name>ovirtmgmt</name>
        </network>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('network', network, types.Network),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(network, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List logical networks.
        For example:
        [source]
        ----
        GET /ovirt-engine/api/networks
        ----
        Will respond:
        [source,xml]
        ----
        <networks>
          <network href="/ovirt-engine/api/networks/123" id="123">
            <name>ovirtmgmt</name>
            <description>Default Management Network</description>
            <link href="/ovirt-engine/api/networks/123/permissions" rel="permissions"/>
            <link href="/ovirt-engine/api/networks/123/vnicprofiles" rel="vnicprofiles"/>
            <link href="/ovirt-engine/api/networks/123/networklabels" rel="networklabels"/>
            <mtu>0</mtu>
            <stp>false</stp>
            <usages>
              <usage>vm</usage>
            </usages>
            <data_center href="/ovirt-engine/api/datacenters/456" id="456"/>
          </network>
          ...
        </networks>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.

        `search`:: A query string used to restrict the returned networks.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def network_service(self, id):
        """
        Reference to the service that manages a specific network.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return NetworkService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.network_service(path)
        return self.network_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'NetworksService:%s' % self._path


class OpenstackImageService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackImageService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def import_(
        self,
        async=None,
        cluster=None,
        disk=None,
        import_as_template=None,
        storage_domain=None,
        template=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Imports a virtual machine from a Glance image storage domain.
        For example, to import the image with identifier `456` from the
        storage domain with identifier `123` send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/openstackimageproviders/123/images/456/import
        ----
        With a request body like this:
        [source,xml]
        ----
        <action>
          <storage_domain>
            <name>images0</name>
          </storage_domain>
          <cluster>
            <name>images0</name>
          </cluster>
        </action>
        ----


        This method supports the following parameters:

        `import_as_template`:: Indicates whether the image should be imported as a template.

        `cluster`:: This parameter is mandatory in case of using `import_as_template` and indicates which cluster should be used
        for import glance image as template.

        `async`:: Indicates if the import should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('cluster', cluster, types.Cluster),
            ('disk', disk, types.Disk),
            ('import_as_template', import_as_template, bool),
            ('storage_domain', storage_domain, types.StorageDomain),
            ('template', template, types.Template),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            cluster=cluster,
            disk=disk,
            import_as_template=import_as_template,
            storage_domain=storage_domain,
            template=template,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'import', None, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'OpenstackImageService:%s' % self._path


class OpenstackImageProviderService(ExternalProviderService):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackImageProviderService, self).__init__(connection, path)
        self._certificates_service = None
        self._images_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def import_certificates(
        self,
        certificates=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('certificates', certificates, list),
        ])

        # Populate the action:
        action = types.Action(
            certificates=certificates,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'importcertificates', None, headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def test_connectivity(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the test should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'testconnectivity', None, headers, query, wait)

    def update(
        self,
        provider,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('provider', provider, types.OpenStackImageProvider),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(provider, headers, query, wait)

    def certificates_service(self):
        """
        """
        return ExternalProviderCertificatesService(self._connection, '%s/certificates' % self._path)

    def images_service(self):
        """
        """
        return OpenstackImagesService(self._connection, '%s/images' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'certificates':
            return self.certificates_service()
        if path.startswith('certificates/'):
            return self.certificates_service().service(path[13:])
        if path == 'images':
            return self.images_service()
        if path.startswith('images/'):
            return self.images_service().service(path[7:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'OpenstackImageProviderService:%s' % self._path


class OpenstackImageProvidersService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackImageProvidersService, self).__init__(connection, path)
        self._provider_service = None

    def add(
        self,
        provider,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('provider', provider, types.OpenStackImageProvider),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(provider, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of providers to return. If not specified all the providers are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def provider_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return OpenstackImageProviderService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.provider_service(path)
        return self.provider_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'OpenstackImageProvidersService:%s' % self._path


class OpenstackImagesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackImagesService, self).__init__(connection, path)
        self._image_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists the images of a Glance image storage domain.


        This method supports the following parameters:

        `max`:: Sets the maximum number of images to return. If not specified all the images are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def image_service(self, id):
        """
        Returns a reference to the service that manages a specific image.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return OpenstackImageService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.image_service(path)
        return self.image_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'OpenstackImagesService:%s' % self._path


class OpenstackNetworkService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackNetworkService, self).__init__(connection, path)
        self._subnets_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def import_(
        self,
        async=None,
        data_center=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation imports an external network into oVirt.
        The network will be added to the data center specified.


        This method supports the following parameters:

        `data_center`:: The data center into which the network is to be imported.
        Data center is mandatory, and can be specified
        using the `id` or `name` attributes, the rest of
        the attributes will be ignored.

        `async`:: Indicates if the import should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('data_center', data_center, types.DataCenter),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            data_center=data_center,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'import', None, headers, query, wait)

    def subnets_service(self):
        """
        """
        return OpenstackSubnetsService(self._connection, '%s/subnets' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'subnets':
            return self.subnets_service()
        if path.startswith('subnets/'):
            return self.subnets_service().service(path[8:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'OpenstackNetworkService:%s' % self._path


class OpenstackNetworkProviderService(ExternalProviderService):
    """
    This service manages OpenStack network provider.

    """

    def __init__(self, connection, path):
        super(OpenstackNetworkProviderService, self).__init__(connection, path)
        self._certificates_service = None
        self._networks_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns the representation of the object managed by this service.
        For example, to get the OpenStack network provider with identifier `1234`, send a request like this:
        [source]
        ----
        GET /ovirt-engine/api/openstacknetworkproviders/1234
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def import_certificates(
        self,
        certificates=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('certificates', certificates, list),
        ])

        # Populate the action:
        action = types.Action(
            certificates=certificates,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'importcertificates', None, headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the provider.
        For example, to remove the OpenStack network provider with identifier `1234`, send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/openstacknetworkproviders/1234
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def test_connectivity(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the test should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'testconnectivity', None, headers, query, wait)

    def update(
        self,
        provider,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the provider.
        For example, to update `provider_name`, `requires_authentication`, `url`, `tenant_name` and `type` properties,
        for the OpenStack network provider with identifier `1234`, send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/openstacknetworkproviders/1234
        ----
        With a request body like this:
        [source,xml]
        ----
        <openstack_network_provider>
          <name>ovn-network-provider</name>
          <requires_authentication>false</requires_authentication>
          <url>http://some_server_url.domain.com:9696</url>
          <tenant_name>oVirt</tenant_name>
          <type>external</type>
        </openstack_network_provider>
        ----


        This method supports the following parameters:

        `provider`:: The provider to update.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('provider', provider, types.OpenStackNetworkProvider),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(provider, headers, query, wait)

    def certificates_service(self):
        """
        """
        return ExternalProviderCertificatesService(self._connection, '%s/certificates' % self._path)

    def networks_service(self):
        """
        Reference to OpenStack networks service.

        """
        return OpenstackNetworksService(self._connection, '%s/networks' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'certificates':
            return self.certificates_service()
        if path.startswith('certificates/'):
            return self.certificates_service().service(path[13:])
        if path == 'networks':
            return self.networks_service()
        if path.startswith('networks/'):
            return self.networks_service().service(path[9:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'OpenstackNetworkProviderService:%s' % self._path


class OpenstackNetworkProvidersService(Service):
    """
    This service manages OpenStack network providers.

    """

    def __init__(self, connection, path):
        super(OpenstackNetworkProvidersService, self).__init__(connection, path)
        self._provider_service = None

    def add(
        self,
        provider,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        The operation adds a new network provider to the system.
        If the `type` property is not present, a default value of `NEUTRON` will be used.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('provider', provider, types.OpenStackNetworkProvider),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(provider, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of providers to return. If not specified all the providers are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def provider_service(self, id):
        """
        Reference to OpenStack network provider service.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return OpenstackNetworkProviderService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.provider_service(path)
        return self.provider_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'OpenstackNetworkProvidersService:%s' % self._path


class OpenstackNetworksService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackNetworksService, self).__init__(connection, path)
        self._network_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def network_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return OpenstackNetworkService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.network_service(path)
        return self.network_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'OpenstackNetworksService:%s' % self._path


class OpenstackSubnetService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackSubnetService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'OpenstackSubnetService:%s' % self._path


class OpenstackSubnetsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackSubnetsService, self).__init__(connection, path)
        self._subnet_service = None

    def add(
        self,
        subnet,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('subnet', subnet, types.OpenStackSubnet),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(subnet, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of sub-networks to return. If not specified all the sub-networks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def subnet_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return OpenstackSubnetService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.subnet_service(path)
        return self.subnet_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'OpenstackSubnetsService:%s' % self._path


class OpenstackVolumeAuthenticationKeyService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackVolumeAuthenticationKeyService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        key,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('key', key, types.OpenstackVolumeAuthenticationKey),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_update(key, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'OpenstackVolumeAuthenticationKeyService:%s' % self._path


class OpenstackVolumeAuthenticationKeysService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackVolumeAuthenticationKeysService, self).__init__(connection, path)
        self._key_service = None

    def add(
        self,
        key,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('key', key, types.OpenstackVolumeAuthenticationKey),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(key, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of keys to return. If not specified all the keys are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def key_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return OpenstackVolumeAuthenticationKeyService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.key_service(path)
        return self.key_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'OpenstackVolumeAuthenticationKeysService:%s' % self._path


class OpenstackVolumeProviderService(ExternalProviderService):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackVolumeProviderService, self).__init__(connection, path)
        self._authentication_keys_service = None
        self._certificates_service = None
        self._volume_types_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def import_certificates(
        self,
        certificates=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('certificates', certificates, list),
        ])

        # Populate the action:
        action = types.Action(
            certificates=certificates,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'importcertificates', None, headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def test_connectivity(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the test should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'testconnectivity', None, headers, query, wait)

    def update(
        self,
        provider,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('provider', provider, types.OpenStackVolumeProvider),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(provider, headers, query, wait)

    def authentication_keys_service(self):
        """
        """
        return OpenstackVolumeAuthenticationKeysService(self._connection, '%s/authenticationkeys' % self._path)

    def certificates_service(self):
        """
        """
        return ExternalProviderCertificatesService(self._connection, '%s/certificates' % self._path)

    def volume_types_service(self):
        """
        """
        return OpenstackVolumeTypesService(self._connection, '%s/volumetypes' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'authenticationkeys':
            return self.authentication_keys_service()
        if path.startswith('authenticationkeys/'):
            return self.authentication_keys_service().service(path[19:])
        if path == 'certificates':
            return self.certificates_service()
        if path.startswith('certificates/'):
            return self.certificates_service().service(path[13:])
        if path == 'volumetypes':
            return self.volume_types_service()
        if path.startswith('volumetypes/'):
            return self.volume_types_service().service(path[12:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'OpenstackVolumeProviderService:%s' % self._path


class OpenstackVolumeProvidersService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackVolumeProvidersService, self).__init__(connection, path)
        self._provider_service = None

    def add(
        self,
        provider,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adds a new volume provider.
        For example:
        [source]
        ----
        POST /ovirt-engine/api/openstackvolumeproviders
        ----
        With a request body like this:
        [source,xml]
        ----
        <openstack_volume_provider>
          <name>mycinder</name>
          <url>https://mycinder.example.com:8776</url>
          <data_center>
            <name>mydc</name>
          </data_center>
          <requires_authentication>true</requires_authentication>
          <username>admin</username>
          <password>mypassword</password>
          <tenant_name>mytenant</tenant_name>
        </openstack_volume_provider>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('provider', provider, types.OpenStackVolumeProvider),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(provider, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the list of volume providers.


        This method supports the following parameters:

        `max`:: Sets the maximum number of providers to return. If not specified all the providers are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def provider_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return OpenstackVolumeProviderService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.provider_service(path)
        return self.provider_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'OpenstackVolumeProvidersService:%s' % self._path


class OpenstackVolumeTypeService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackVolumeTypeService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'OpenstackVolumeTypeService:%s' % self._path


class OpenstackVolumeTypesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OpenstackVolumeTypesService, self).__init__(connection, path)
        self._type_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of volume types to return. If not specified all the volume types are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def type_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return OpenstackVolumeTypeService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.type_service(path)
        return self.type_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'OpenstackVolumeTypesService:%s' % self._path


class OperatingSystemService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OperatingSystemService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'OperatingSystemService:%s' % self._path


class OperatingSystemsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(OperatingSystemsService, self).__init__(connection, path)
        self._operating_system_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def operating_system_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return OperatingSystemService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.operating_system_service(path)
        return self.operating_system_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'OperatingSystemsService:%s' % self._path


class PermissionService(Service):
    """
    """

    def __init__(self, connection, path):
        super(PermissionService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'PermissionService:%s' % self._path


class PermitService(Service):
    """
    A service to manage a specific permit of the role.

    """

    def __init__(self, connection, path):
        super(PermitService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets the information about the permit of the role.
        For example to retrieve the information about the permit with the id `456` of the role with the id `123`
        send a request like this:
        ....
        GET /ovirt-engine/api/roles/123/permits/456
        ....
        [source,xml]
        ----
        <permit href="/ovirt-engine/api/roles/123/permits/456" id="456">
          <name>change_vm_cd</name>
          <administrative>false</administrative>
          <role href="/ovirt-engine/api/roles/123" id="123"/>
        </permit>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the permit from the role.
        For example to remove the permit with id `456` from the role with id `123` send a request like this:
        ....
        DELETE /ovirt-engine/api/roles/123/permits/456
        ....


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'PermitService:%s' % self._path


class PermitsService(Service):
    """
    Represents a permits sub-collection of the specific role.

    """

    def __init__(self, connection, path):
        super(PermitsService, self).__init__(connection, path)
        self._permit_service = None

    def add(
        self,
        permit,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adds a permit to the role. The permit name can be retrieved from the <<services/cluster_levels>> service.
        For example to assign a permit `create_vm` to the role with id `123` send a request like this:
        ....
        POST /ovirt-engine/api/roles/123/permits
        ....
        With a request body like this:
        [source,xml]
        ----
        <permit>
          <name>create_vm</name>
        </permit>
        ----


        This method supports the following parameters:

        `permit`:: The permit to add.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('permit', permit, types.Permit),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(permit, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List the permits of the role.
        For example to list the permits of the role with the id `123` send a request like this:
        ....
        GET /ovirt-engine/api/roles/123/permits
        ....
        [source,xml]
        ----
        <permits>
          <permit href="/ovirt-engine/api/roles/123/permits/5" id="5">
            <name>change_vm_cd</name>
            <administrative>false</administrative>
            <role href="/ovirt-engine/api/roles/123" id="123"/>
          </permit>
          <permit href="/ovirt-engine/api/roles/123/permits/7" id="7">
            <name>connect_to_vm</name>
            <administrative>false</administrative>
            <role href="/ovirt-engine/api/roles/123" id="123"/>
          </permit>
        </permits>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of permits to return. If not specified all the permits are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def permit_service(self, id):
        """
        Sub-resource locator method, returns individual permit resource on which the remainder of the URI is dispatched.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return PermitService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.permit_service(path)
        return self.permit_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'PermitsService:%s' % self._path


class QosService(Service):
    """
    """

    def __init__(self, connection, path):
        super(QosService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        qos,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('qos', qos, types.Qos),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(qos, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'QosService:%s' % self._path


class QossService(Service):
    """
    """

    def __init__(self, connection, path):
        super(QossService, self).__init__(connection, path)
        self._qos_service = None

    def add(
        self,
        qos,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('qos', qos, types.Qos),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(qos, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of QoS descriptors to return. If not specified all the descriptors are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def qos_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return QosService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.qos_service(path)
        return self.qos_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'QossService:%s' % self._path


class QuotaService(Service):
    """
    """

    def __init__(self, connection, path):
        super(QuotaService, self).__init__(connection, path)
        self._permissions_service = None
        self._quota_cluster_limits_service = None
        self._quota_storage_limits_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves a quota.
        An example of retrieving a quota:
        [source]
        ----
        GET /ovirt-engine/api/datacenters/123/quotas/456
        ----
        [source,xml]
        ----
        <quota id="456">
          <name>myquota</name>
          <description>My new quota for virtual machines</description>
          <cluster_hard_limit_pct>20</cluster_hard_limit_pct>
          <cluster_soft_limit_pct>80</cluster_soft_limit_pct>
          <storage_hard_limit_pct>20</storage_hard_limit_pct>
          <storage_soft_limit_pct>80</storage_soft_limit_pct>
        </quota>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Delete a quota.
        An example of deleting a quota:
        [source]
        ----
        DELETE /ovirt-engine/api/datacenters/123-456/quotas/654-321
        -0472718ab224 HTTP/1.1
        Accept: application/xml
        Content-type: application/xml
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        quota,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates a quota.
        An example of updating a quota:
        [source]
        ----
        PUT /ovirt-engine/api/datacenters/123/quotas/456
        ----
        [source,xml]
        ----
        <quota>
          <cluster_hard_limit_pct>30</cluster_hard_limit_pct>
          <cluster_soft_limit_pct>70</cluster_soft_limit_pct>
          <storage_hard_limit_pct>20</storage_hard_limit_pct>
          <storage_soft_limit_pct>80</storage_soft_limit_pct>
        </quota>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('quota', quota, types.Quota),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(quota, headers, query, wait)

    def permissions_service(self):
        """
        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def quota_cluster_limits_service(self):
        """
        """
        return QuotaClusterLimitsService(self._connection, '%s/quotaclusterlimits' % self._path)

    def quota_storage_limits_service(self):
        """
        """
        return QuotaStorageLimitsService(self._connection, '%s/quotastoragelimits' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'quotaclusterlimits':
            return self.quota_cluster_limits_service()
        if path.startswith('quotaclusterlimits/'):
            return self.quota_cluster_limits_service().service(path[19:])
        if path == 'quotastoragelimits':
            return self.quota_storage_limits_service()
        if path.startswith('quotastoragelimits/'):
            return self.quota_storage_limits_service().service(path[19:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'QuotaService:%s' % self._path


class QuotaClusterLimitService(Service):
    """
    """

    def __init__(self, connection, path):
        super(QuotaClusterLimitService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'QuotaClusterLimitService:%s' % self._path


class QuotaClusterLimitsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(QuotaClusterLimitsService, self).__init__(connection, path)
        self._limit_service = None

    def add(
        self,
        limit,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('limit', limit, types.QuotaClusterLimit),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(limit, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of limits to return. If not specified all the limits are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def limit_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return QuotaClusterLimitService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.limit_service(path)
        return self.limit_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'QuotaClusterLimitsService:%s' % self._path


class QuotaStorageLimitService(Service):
    """
    """

    def __init__(self, connection, path):
        super(QuotaStorageLimitService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the update should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'QuotaStorageLimitService:%s' % self._path


class QuotaStorageLimitsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(QuotaStorageLimitsService, self).__init__(connection, path)
        self._limit_service = None

    def add(
        self,
        limit,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('limit', limit, types.QuotaStorageLimit),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(limit, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of limits to return. If not specified all the limits are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def limit_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return QuotaStorageLimitService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.limit_service(path)
        return self.limit_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'QuotaStorageLimitsService:%s' % self._path


class QuotasService(Service):
    """
    """

    def __init__(self, connection, path):
        super(QuotasService, self).__init__(connection, path)
        self._quota_service = None

    def add(
        self,
        quota,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new quota.
        An example of creating a new quota:
        [source]
        ----
        POST /ovirt-engine/api/datacenters/123/quotas
        ----
        [source,xml]
        ----
        <quota>
          <name>myquota</name>
          <description>My new quota for virtual machines</description>
        </quota>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('quota', quota, types.Quota),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(quota, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists quotas of a data center


        This method supports the following parameters:

        `max`:: Sets the maximum number of quota descriptors to return. If not specified all the descriptors are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def quota_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return QuotaService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.quota_service(path)
        return self.quota_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'QuotasService:%s' % self._path


class RoleService(Service):
    """
    """

    def __init__(self, connection, path):
        super(RoleService, self).__init__(connection, path)
        self._permits_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get the role.
        [source]
        ----
        GET /ovirt-engine/api/roles/123
        ----
        You will receive XML response like this one:
        [source,xml]
        ----
        <role id="123">
          <name>MyRole</name>
          <description>MyRole description</description>
          <link href="/ovirt-engine/api/roles/123/permits" rel="permits"/>
          <administrative>true</administrative>
          <mutable>false</mutable>
        </role>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the role.
        To remove the role you need to know its id, then send request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/roles/{role_id}
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        role,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates a role. You are allowed to update `name`, `description` and `administrative` attributes after role is
        created. Within this endpoint you can't add or remove roles permits you need to use
        <<services/permits, service>> that manages permits of role.
        For example to update role's `name`, `description` and `administrative` attributes send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/roles/123
        ----
        With a request body like this:
        [source,xml]
        ----
        <role>
          <name>MyNewRoleName</name>
          <description>My new description of the role</description>
          <administrative>true</administrative>
        </group>
        ----


        This method supports the following parameters:

        `role`:: Updated role.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('role', role, types.Role),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(role, headers, query, wait)

    def permits_service(self):
        """
        Sub-resource locator method, returns permits service.

        """
        return PermitsService(self._connection, '%s/permits' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permits':
            return self.permits_service()
        if path.startswith('permits/'):
            return self.permits_service().service(path[8:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'RoleService:%s' % self._path


class RolesService(Service):
    """
    Provides read-only access to the global set of roles

    """

    def __init__(self, connection, path):
        super(RolesService, self).__init__(connection, path)
        self._role_service = None

    def add(
        self,
        role,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Create a new role. The role can be administrative or non-administrative and can have different permits.
        For example, to add the `MyRole` non-administrative role with permits to login and create virtual machines
        send a request like this (note that you have to pass permit id):
        [source]
        ----
        POST /ovirt-engine/api/roles
        ----
        With a request body like this:
        [source,xml]
        ----
        <role>
          <name>MyRole</name>
          <description>My custom role to create virtual machines</description>
          <administrative>false</administrative>
          <permits>
            <permit id="1"/>
            <permit id="1300"/>
          </permits>
        </group>
        ----


        This method supports the following parameters:

        `role`:: Role that will be added.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('role', role, types.Role),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(role, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List roles.
        [source]
        ----
        GET /ovirt-engine/api/roles
        ----
        You will receive response in XML like this one:
        [source,xml]
        ----
        <roles>
          <role id="123">
             <name>SuperUser</name>
             <description>Roles management administrator</description>
             <link href="/ovirt-engine/api/roles/123/permits" rel="permits"/>
             <administrative>true</administrative>
             <mutable>false</mutable>
          </role>
          ...
        </roles>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of roles to return. If not specified all the roles are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def role_service(self, id):
        """
        Sub-resource locator method, returns individual role resource on which the remainder of the URI is dispatched.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return RoleService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.role_service(path)
        return self.role_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'RolesService:%s' % self._path


class SchedulingPoliciesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SchedulingPoliciesService, self).__init__(connection, path)
        self._policy_service = None

    def add(
        self,
        policy,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('policy', policy, types.SchedulingPolicy),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(policy, headers, query, wait)

    def list(
        self,
        filter=None,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of policies to return. If not specified all the policies are returned.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def policy_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return SchedulingPolicyService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.policy_service(path)
        return self.policy_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'SchedulingPoliciesService:%s' % self._path


class SchedulingPolicyService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SchedulingPolicyService, self).__init__(connection, path)
        self._balances_service = None
        self._filters_service = None
        self._weights_service = None

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        policy,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('policy', policy, types.SchedulingPolicy),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(policy, headers, query, wait)

    def balances_service(self):
        """
        """
        return BalancesService(self._connection, '%s/balances' % self._path)

    def filters_service(self):
        """
        """
        return FiltersService(self._connection, '%s/filters' % self._path)

    def weights_service(self):
        """
        """
        return WeightsService(self._connection, '%s/weights' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'balances':
            return self.balances_service()
        if path.startswith('balances/'):
            return self.balances_service().service(path[9:])
        if path == 'filters':
            return self.filters_service()
        if path.startswith('filters/'):
            return self.filters_service().service(path[8:])
        if path == 'weights':
            return self.weights_service()
        if path.startswith('weights/'):
            return self.weights_service().service(path[8:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'SchedulingPolicyService:%s' % self._path


class SchedulingPolicyUnitService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SchedulingPolicyUnitService, self).__init__(connection, path)

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'SchedulingPolicyUnitService:%s' % self._path


class SchedulingPolicyUnitsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SchedulingPolicyUnitsService, self).__init__(connection, path)
        self._unit_service = None

    def list(
        self,
        filter=None,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of policy units to return. If not specified all the policy units are returned.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def unit_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return SchedulingPolicyUnitService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.unit_service(path)
        return self.unit_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'SchedulingPolicyUnitsService:%s' % self._path


class SnapshotService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SnapshotService, self).__init__(connection, path)
        self._cdroms_service = None
        self._disks_service = None
        self._nics_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        all_content=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `all_content`:: Indicates if all the attributes of the virtual machine snapshot should be included in the response.
        By default the attribute `initialization.configuration.data` is excluded.
        For example, to retrieve the complete representation of the snapshot with id `456` of the virtual machine
        with id `123` send a request like this:
        ....
        GET /ovirt-engine/api/vms/123/snapshots/456?all_content=true
        ....

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('all_content', all_content, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async
        if all_content is not None:
            all_content = Writer.render_boolean(all_content)
            query['all_content'] = all_content

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def restore(
        self,
        async=None,
        disks=None,
        restore_memory=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Restores a virtual machine snapshot.
        For example, to restore the snapshot with identifier `456` of virtual machine with identifier `123` send a
        request like this:
        [source]
        ----
        POST /ovirt-engine/api/vms/123/snapshots/456/restore
        ----
        With an empty `action` in the body:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the restore should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('disks', disks, list),
            ('restore_memory', restore_memory, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            disks=disks,
            restore_memory=restore_memory,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'restore', None, headers, query, wait)

    def cdroms_service(self):
        """
        """
        return SnapshotCdromsService(self._connection, '%s/cdroms' % self._path)

    def disks_service(self):
        """
        """
        return SnapshotDisksService(self._connection, '%s/disks' % self._path)

    def nics_service(self):
        """
        """
        return SnapshotNicsService(self._connection, '%s/nics' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'cdroms':
            return self.cdroms_service()
        if path.startswith('cdroms/'):
            return self.cdroms_service().service(path[7:])
        if path == 'disks':
            return self.disks_service()
        if path.startswith('disks/'):
            return self.disks_service().service(path[6:])
        if path == 'nics':
            return self.nics_service()
        if path.startswith('nics/'):
            return self.nics_service().service(path[5:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'SnapshotService:%s' % self._path


class SnapshotCdromService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SnapshotCdromService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'SnapshotCdromService:%s' % self._path


class SnapshotCdromsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SnapshotCdromsService, self).__init__(connection, path)
        self._cdrom_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of CDROMS to return. If not specified all the CDROMS are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def cdrom_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return SnapshotCdromService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.cdrom_service(path)
        return self.cdrom_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'SnapshotCdromsService:%s' % self._path


class SnapshotDiskService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SnapshotDiskService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'SnapshotDiskService:%s' % self._path


class SnapshotDisksService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SnapshotDisksService, self).__init__(connection, path)
        self._disk_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def disk_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return SnapshotDiskService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.disk_service(path)
        return self.disk_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'SnapshotDisksService:%s' % self._path


class SnapshotNicService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SnapshotNicService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'SnapshotNicService:%s' % self._path


class SnapshotNicsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SnapshotNicsService, self).__init__(connection, path)
        self._nic_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of NICs to return. If not specified all the NICs are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def nic_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return SnapshotNicService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.nic_service(path)
        return self.nic_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'SnapshotNicsService:%s' % self._path


class SnapshotsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SnapshotsService, self).__init__(connection, path)
        self._snapshot_service = None

    def add(
        self,
        snapshot,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a virtual machine snapshot.
        For example, to create a new snapshot for virtual machine `123` send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/vms/123/snapshots
        ----
        With a request body like this:
        [source,xml]
        ----
        <snapshot>
          <description>My snapshot</description>
        </snapshot>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('snapshot', snapshot, types.Snapshot),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(snapshot, headers, query, wait)

    def list(
        self,
        all_content=None,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of snapshots to return. If not specified all the snapshots are returned.

        `all_content`:: Indicates if all the attributes of the virtual machine snapshot should be included in the response.
        By default the attribute `initialization.configuration.data` is excluded.
        For example, to retrieve the complete representation of the virtual machine with id `123` snapshots send a
        request like this:
        ....
        GET /ovirt-engine/api/vms/123/snapshots?all_content=true
        ....

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('all_content', all_content, bool),
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if all_content is not None:
            all_content = Writer.render_boolean(all_content)
            query['all_content'] = all_content
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def snapshot_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return SnapshotService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.snapshot_service(path)
        return self.snapshot_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'SnapshotsService:%s' % self._path


class SshPublicKeyService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SshPublicKeyService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        key,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('key', key, types.SshPublicKey),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(key, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'SshPublicKeyService:%s' % self._path


class SshPublicKeysService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SshPublicKeysService, self).__init__(connection, path)
        self._key_service = None

    def add(
        self,
        key,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('key', key, types.SshPublicKey),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(key, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of keys to return. If not specified all the keys are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def key_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return SshPublicKeyService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.key_service(path)
        return self.key_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'SshPublicKeysService:%s' % self._path


class StatisticService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StatisticService, self).__init__(connection, path)

    def get(
        self,
        statistic=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('statistic', statistic, types.Statistic),
        ])

        # Build the URL:
        query = query or {}
        if statistic is not None:
            query['statistic'] = statistic

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StatisticService:%s' % self._path


class StatisticsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StatisticsService, self).__init__(connection, path)
        self._statistic_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves a list of statistics.
        For example, to retrieve the statistics for virtual machine `123` send a
        request like this:
        [source]
        ----
        GET /ovirt-engine/api/vms/123/statistics
        ----
        The result will be like this:
        [source,xml]
        ----
        <statistics>
          <statistic href="/ovirt-engine/api/vms/123/statistics/456" id="456">
            <name>memory.installed</name>
            <description>Total memory configured</description>
            <kind>gauge</kind>
            <type>integer</type>
            <unit>bytes</unit>
            <values>
              <value>
                <datum>1073741824</datum>
              </value>
            </values>
            <vm href="/ovirt-engine/api/vms/123" id="123"/>
          </statistic>
          ...
        </statistics>
        ----
        Just a single part of the statistics can be retrieved by specifying its id at the end of the URI. That means:
        [source]
        ----
        GET /ovirt-engine/api/vms/123/statistics/456
        ----
        Outputs:
        [source,xml]
        ----
        <statistic href="/ovirt-engine/api/vms/123/statistics/456" id="456">
          <name>memory.installed</name>
          <description>Total memory configured</description>
          <kind>gauge</kind>
          <type>integer</type>
          <unit>bytes</unit>
          <values>
            <value>
              <datum>1073741824</datum>
            </value>
          </values>
          <vm href="/ovirt-engine/api/vms/123" id="123"/>
        </statistic>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of statistics to return. If not specified all the statistics are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def statistic_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return StatisticService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.statistic_service(path)
        return self.statistic_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'StatisticsService:%s' % self._path


class StepService(MeasurableService):
    """
    A service to manage a step.

    """

    def __init__(self, connection, path):
        super(StepService, self).__init__(connection, path)
        self._statistics_service = None

    def end(
        self,
        async=None,
        force=None,
        succeeded=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Marks an external step execution as ended.
        For example, to terminate a step with identifier `456` which belongs to a `job` with identifier `123` send the
        following request:
        [source]
        ----
        POST /ovirt-engine/api/jobs/123/steps/456/end
        ----
        With the following request body:
        [source,xml]
        ----
        <action>
          <force>true</force>
          <succeeded>true</succeeded>
        </action>
        ----


        This method supports the following parameters:

        `force`:: Indicates if the step should be forcibly terminated.

        `succeeded`:: Indicates if the step should be marked as successfully finished or as failed.
        This parameter is optional, and the default value is `true`.

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('force', force, bool),
            ('succeeded', succeeded, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            force=force,
            succeeded=succeeded,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'end', None, headers, query, wait)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves a step.
        [source]
        ----
        GET /ovirt-engine/api/jobs/123/steps/456
        ----
        You will receive response in XML like this one:
        [source,xml]
        ----
        <step href="/ovirt-engine/api/jobs/123/steps/456" id="456">
          <actions>
            <link href="/ovirt-engine/api/jobs/123/steps/456/end" rel="end"/>
          </actions>
          <description>Validating</description>
          <end_time>2016-12-12T23:07:26.627+02:00</end_time>
          <external>false</external>
          <number>0</number>
          <start_time>2016-12-12T23:07:26.605+02:00</start_time>
          <status>finished</status>
          <type>validating</type>
          <job href="/ovirt-engine/api/jobs/123" id="123"/>
        </step>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StepService:%s' % self._path


class StepsService(Service):
    """
    A service to manage steps.

    """

    def __init__(self, connection, path):
        super(StepsService, self).__init__(connection, path)
        self._step_service = None

    def add(
        self,
        step,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add an external step to an existing job or to an existing step.
        For example, to add a step to `job` with identifier `123` send the
        following request:
        [source]
        ----
        POST /ovirt-engine/api/jobs/123/steps
        ----
        With the following request body:
        [source,xml]
        ----
        <step>
          <description>Validating</description>
          <start_time>2016-12-12T23:07:26.605+02:00</start_time>
          <status>started</status>
          <type>validating</type>
        </step>
        ----
        The response should look like:
        [source,xml]
        ----
        <step href="/ovirt-engine/api/jobs/123/steps/456" id="456">
          <actions>
            <link href="/ovirt-engine/api/jobs/123/steps/456/end" rel="end"/>
          </actions>
          <description>Validating</description>
          <link href="/ovirt-engine/api/jobs/123/steps/456/statistics" rel="statistics"/>
          <external>true</external>
          <number>2</number>
          <start_time>2016-12-13T01:06:15.380+02:00</start_time>
          <status>started</status>
          <type>validating</type>
          <job href="/ovirt-engine/api/jobs/123" id="123"/>
        </step>
        ----


        This method supports the following parameters:

        `step`:: Step that will be added.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('step', step, types.Step),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(step, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the representation of the steps.
        [source]
        ----
        GET /ovirt-engine/api/job/123/steps
        ----
        You will receive response in XML like this one:
        [source,xml]
        ----
        <steps>
          <step href="/ovirt-engine/api/jobs/123/steps/456" id="456">
            <actions>
              <link href="/ovirt-engine/api/jobs/123/steps/456/end" rel="end"/>
            </actions>
            <description>Validating</description>
            <link href="/ovirt-engine/api/jobs/123/steps/456/statistics" rel="statistics"/>
            <external>true</external>
            <number>2</number>
            <start_time>2016-12-13T01:06:15.380+02:00</start_time>
            <status>started</status>
            <type>validating</type>
            <job href="/ovirt-engine/api/jobs/123" id="123"/>
          </step>
          ...
        </steps>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of steps to return. If not specified all the steps are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def step_service(self, id):
        """
        Reference to the step service.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return StepService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.step_service(path)
        return self.step_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'StepsService:%s' % self._path


class StorageService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageService, self).__init__(connection, path)

    def get(
        self,
        report_status=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `report_status`:: Indicates if the status of the LUNs in the storage should be checked.
        Checking the status of the LUN is an heavy weight operation and
        this data is not always needed by the user.
        This parameter will give the option to not perform the status check of the LUNs.
        The default is `true` for backward compatibility.
        Here an example with the LUN status :
        [source,xml]
        ----
        <host_storage id="360014051136c20574f743bdbd28177fd">
          <logical_units>
            <logical_unit id="360014051136c20574f743bdbd28177fd">
              <lun_mapping>0</lun_mapping>
              <paths>1</paths>
              <product_id>lun0</product_id>
              <serial>SLIO-ORG_lun0_1136c205-74f7-43bd-bd28-177fd5ce6993</serial>
              <size>10737418240</size>
              <status>used</status>
              <vendor_id>LIO-ORG</vendor_id>
              <volume_group_id>O9Du7I-RahN-ECe1-dZ1w-nh0b-64io-MNzIBZ</volume_group_id>
            </logical_unit>
          </logical_units>
          <type>iscsi</type>
          <host id="8bb5ade5-e988-4000-8b93-dbfc6717fe50"/>
        </host_storage>
        ----
        Here an example without the LUN status :
        [source,xml]
        ----
        <host_storage id="360014051136c20574f743bdbd28177fd">
          <logical_units>
            <logical_unit id="360014051136c20574f743bdbd28177fd">
              <lun_mapping>0</lun_mapping>
              <paths>1</paths>
              <product_id>lun0</product_id>
              <serial>SLIO-ORG_lun0_1136c205-74f7-43bd-bd28-177fd5ce6993</serial>
              <size>10737418240</size>
              <vendor_id>LIO-ORG</vendor_id>
              <volume_group_id>O9Du7I-RahN-ECe1-dZ1w-nh0b-64io-MNzIBZ</volume_group_id>
            </logical_unit>
          </logical_units>
          <type>iscsi</type>
          <host id="8bb5ade5-e988-4000-8b93-dbfc6717fe50"/>
        </host_storage>
        ----

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('report_status', report_status, bool),
        ])

        # Build the URL:
        query = query or {}
        if report_status is not None:
            report_status = Writer.render_boolean(report_status)
            query['report_status'] = report_status

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StorageService:%s' % self._path


class StorageDomainService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageDomainService, self).__init__(connection, path)
        self._disk_profiles_service = None
        self._disk_snapshots_service = None
        self._disks_service = None
        self._files_service = None
        self._images_service = None
        self._permissions_service = None
        self._storage_connections_service = None
        self._templates_service = None
        self._vms_service = None

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def is_attached(
        self,
        async=None,
        host=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('host', host, types.Host),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            host=host,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'isattached', 'is_attached', headers, query, wait)

    def reduce_luns(
        self,
        logical_units=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation reduces logical units from the storage domain.
        In order to do so the data stored on the provided logical units will be moved to other logical units of the
        storage domain and only then they will be reduced from the storage domain.
        For example, in order to reduce two logical units from a storage domain send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/storagedomains/123/reduceluns
        ----
        With a request body like this:
        [source,xml]
        ----
         <action>
           <logical_units>
             <logical_unit id="1IET_00010001"/>
             <logical_unit id="1IET_00010002"/>
           </logical_units>
         </action>
        ----


        This method supports the following parameters:

        `logical_units`:: The logical units that needs to be reduced from the storage domain.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('logical_units', logical_units, list),
        ])

        # Populate the action:
        action = types.Action(
            logical_units=logical_units,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'reduceluns', None, headers, query, wait)

    def refresh_luns(
        self,
        async=None,
        logical_units=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation refreshes the LUN size.
        After increasing the size of the underlying LUN on the storage server,
        the user can refresh the LUN size.
        This action forces a rescan of the provided LUNs and
        updates the database with the new size if required.
        For example, in order to refresh the size of two LUNs send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/storagedomains/262b056b-aede-40f1-9666-b883eff59d40/refreshluns
        ----
        With a request body like this:
        [source,xml]
        ----
         <action>
           <logical_units>
             <logical_unit id="1IET_00010001"/>
             <logical_unit id="1IET_00010002"/>
           </logical_units>
         </action>
        ----


        This method supports the following parameters:

        `logical_units`:: The LUNs that need to be refreshed.

        `async`:: Indicates if the refresh should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('logical_units', logical_units, list),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            logical_units=logical_units,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'refreshluns', None, headers, query, wait)

    def remove(
        self,
        host=None,
        format=None,
        destroy=None,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the storage domain.
        Without any special parameters, the storage domain is detached from the system and removed from the database. The
        storage domain can then be imported to the same or different setup, with all the data on it. If the storage isn't
        accessible the operation will fail.
        If the `destroy` parameter is `true` then the operation will always succeed, even if the storage isn't
        accessible, the failure is just ignored and the storage domain is removed from the database anyway.
        If the `format` parameter is `true` then the actual storage is formatted, and the metadata is removed from the
        LUN or directory, so it can no longer be imported to the same or a different setup.


        This method supports the following parameters:

        `host`:: Indicates what host should be used to remove the storage domain.
        This parameter is mandatory, and it can contain the name or the identifier of the host. For example, to use
        the host named `myhost` to remove the storage domain with identifier `123` send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/storagedomains/123?host=myhost
        ----

        `format`:: Indicates if the actual storage should be formatted, removing all the metadata from the underlying LUN or
        directory:
        [source]
        ----
        DELETE /ovirt-engine/api/storagedomains/123?format=true
        ----
        This parameter is optional, and the default value is `false`.

        `destroy`:: Indicates if the operation should succeed, and the storage domain removed from the database, even if the
        storage isn't accessible.
        [source]
        ----
        DELETE /ovirt-engine/api/storagedomains/123?destroy=true
        ----
        This parameter is optional, and the default value is `false`.

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('host', host, str),
            ('format', format, bool),
            ('destroy', destroy, bool),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if host is not None:
            query['host'] = host
        if format is not None:
            format = Writer.render_boolean(format)
            query['format'] = format
        if destroy is not None:
            destroy = Writer.render_boolean(destroy)
            query['destroy'] = destroy
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        storage_domain,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates a storage domain.
        Not all of the <<types/storage_domain,StorageDomain>>'s attributes are updatable post-creation. Those that can be
        updated are: `name`, `description`, `comment`, `warning_low_space_indicator`, `critical_space_action_blocker` and
        `wipe_after_delete` (note that changing the `wipe_after_delete` attribute will not change the wipe after delete
        property of disks that already exist).
        To update the `name` and `wipe_after_delete` attributes of a storage domain with an identifier `123`, send a
        request as follows:
        [source]
        ----
        PUT /ovirt-engine/api/storagedomains/123
        ----
        With a request body as follows:
        [source,xml]
        ----
        <storage_domain>
          <name>data2</name>
          <wipe_after_delete>true</wipe_after_delete>
        </storage_domain>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('storage_domain', storage_domain, types.StorageDomain),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(storage_domain, headers, query, wait)

    def update_ovf_store(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation forces the update of the `OVF_STORE`
        of this storage domain.
        The `OVF_STORE` is a disk image that contains the meta-data
        of virtual machines and disks that reside in the
        storage domain. This meta-data is used in case the
        domain is imported or exported to or from a different
        data center or a different installation.
        By default the `OVF_STORE` is updated periodically
        (set by default to 60 minutes) but users might want to force an
        update after an important change, or when the they believe the
        `OVF_STORE` is corrupt.
        When initiated by the user, `OVF_STORE` update will be performed whether
        an update is needed or not.


        This method supports the following parameters:

        `async`:: Indicates if the `OVF_STORE` update should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'updateovfstore', None, headers, query, wait)

    def disk_profiles_service(self):
        """
        """
        return AssignedDiskProfilesService(self._connection, '%s/diskprofiles' % self._path)

    def disk_snapshots_service(self):
        """
        """
        return DiskSnapshotsService(self._connection, '%s/disksnapshots' % self._path)

    def disks_service(self):
        """
        Reference to the service that manages the disks available in the storage domain.

        """
        return StorageDomainDisksService(self._connection, '%s/disks' % self._path)

    def files_service(self):
        """
        Returns a reference to the service that manages the files available in the storage domain.

        """
        return FilesService(self._connection, '%s/files' % self._path)

    def images_service(self):
        """
        """
        return ImagesService(self._connection, '%s/images' % self._path)

    def permissions_service(self):
        """
        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def storage_connections_service(self):
        """
        Returns a reference to the service that manages the storage connections.

        """
        return StorageDomainServerConnectionsService(self._connection, '%s/storageconnections' % self._path)

    def templates_service(self):
        """
        """
        return StorageDomainTemplatesService(self._connection, '%s/templates' % self._path)

    def vms_service(self):
        """
        """
        return StorageDomainVmsService(self._connection, '%s/vms' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'diskprofiles':
            return self.disk_profiles_service()
        if path.startswith('diskprofiles/'):
            return self.disk_profiles_service().service(path[13:])
        if path == 'disksnapshots':
            return self.disk_snapshots_service()
        if path.startswith('disksnapshots/'):
            return self.disk_snapshots_service().service(path[14:])
        if path == 'disks':
            return self.disks_service()
        if path.startswith('disks/'):
            return self.disks_service().service(path[6:])
        if path == 'files':
            return self.files_service()
        if path.startswith('files/'):
            return self.files_service().service(path[6:])
        if path == 'images':
            return self.images_service()
        if path.startswith('images/'):
            return self.images_service().service(path[7:])
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'storageconnections':
            return self.storage_connections_service()
        if path.startswith('storageconnections/'):
            return self.storage_connections_service().service(path[19:])
        if path == 'templates':
            return self.templates_service()
        if path.startswith('templates/'):
            return self.templates_service().service(path[10:])
        if path == 'vms':
            return self.vms_service()
        if path.startswith('vms/'):
            return self.vms_service().service(path[4:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StorageDomainService:%s' % self._path


class StorageDomainContentDiskService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageDomainContentDiskService, self).__init__(connection, path)

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StorageDomainContentDiskService:%s' % self._path


class StorageDomainContentDisksService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageDomainContentDisksService, self).__init__(connection, path)
        self._disk_service = None

    def list(
        self,
        case_sensitive=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.

        `search`:: A query string used to restrict the returned disks.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def disk_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return StorageDomainContentDiskService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.disk_service(path)
        return self.disk_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'StorageDomainContentDisksService:%s' % self._path


class StorageDomainDiskService(MeasurableService):
    """
    Manages a single disk available in a storage domain.
    IMPORTANT: Since version 4.2 of the engine this service is intended only to list disks available in the storage
    domain, and to register unregistered disks. All the other operations, like copying a disk, moving a disk, etc, have
    been deprecated and will be removed in the future. To perform those operations use the <<services/disks, service
    that manages all the disks of the system>>, or the <<services/disk, service that manages an specific disk>>.

    """

    def __init__(self, connection, path):
        super(StorageDomainDiskService, self).__init__(connection, path)
        self._permissions_service = None
        self._statistics_service = None

    def copy(
        self,
        disk=None,
        storage_domain=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Copies a disk to the specified storage domain.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To copy a disk use the <<services/disk/methods/copy, copy>>
        operation of the service that manages that disk.


        This method supports the following parameters:

        `disk`:: Description of the resulting disk.

        `storage_domain`:: The storage domain where the new disk will be created.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('disk', disk, types.Disk),
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Populate the action:
        action = types.Action(
            disk=disk,
            storage_domain=storage_domain,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'copy', None, headers, query, wait)

    def export(
        self,
        storage_domain=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Exports a disk to an export storage domain.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To export a disk use the <<services/disk/methods/export, export>>
        operation of the service that manages that disk.


        This method supports the following parameters:

        `storage_domain`:: The export storage domain where the disk should be exported to.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Populate the action:
        action = types.Action(
            storage_domain=storage_domain,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'export', None, headers, query, wait)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the description of the disk.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def move(
        self,
        async=None,
        filter=None,
        storage_domain=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Moves a disk to another storage domain.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To move a disk use the <<services/disk/methods/move, move>>
        operation of the service that manages that disk.


        This method supports the following parameters:

        `storage_domain`:: The storage domain where the disk will be moved to.

        `async`:: Indicates if the move should be performed asynchronously.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('filter', filter, bool),
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            filter=filter,
            storage_domain=storage_domain,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'move', None, headers, query, wait)

    def remove(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a disk.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To remove a disk use the <<services/disk/methods/remove, remove>>
        operation of the service that manages that disk.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def sparsify(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Sparsify the disk.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To remove a disk use the <<services/disk/methods/remove, remove>>
        operation of the service that manages that disk.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'sparsify', None, headers, query, wait)

    def update(
        self,
        disk,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the disk.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To update a disk use the
        <<services/disk/methods/update, update>> operation of the service that manages that disk.


        This method supports the following parameters:

        `disk`:: The update to apply to the disk.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('disk', disk, types.Disk),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_update(disk, headers, query, wait)

    def permissions_service(self):
        """
        Reference to the service that manages the permissions assigned to the disk.

        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StorageDomainDiskService:%s' % self._path


class StorageDomainDisksService(Service):
    """
    Manages the collection of disks available inside an specific storage domain.

    """

    def __init__(self, connection, path):
        super(StorageDomainDisksService, self).__init__(connection, path)
        self._disk_service = None

    def add(
        self,
        disk,
        unregistered=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adds or registers a disk.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To add a new disk use the <<services/disks/methods/add, add>>
        operation of the service that manages the disks of the system. To register an unregistered disk use the
        <<services/attached_storage_domain_disk/methods/register, register>> operation of the service that manages
        that disk.


        This method supports the following parameters:

        `disk`:: The disk to add or register.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('disk', disk, types.Disk),
            ('unregistered', unregistered, bool),
        ])

        # Build the URL:
        query = query or {}
        if unregistered is not None:
            unregistered = Writer.render_boolean(unregistered)
            query['unregistered'] = unregistered

        # Send the request and wait for the response:
        return self._internal_add(disk, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieve the list of disks that are available in the storage domain.


        This method supports the following parameters:

        `max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def disk_service(self, id):
        """
        Reference to the service that manages a specific disk.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return StorageDomainDiskService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.disk_service(path)
        return self.disk_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'StorageDomainDisksService:%s' % self._path


class StorageDomainServerConnectionService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageDomainServerConnectionService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Detaches a storage connection from storage.


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StorageDomainServerConnectionService:%s' % self._path


class StorageDomainServerConnectionsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageDomainServerConnectionsService, self).__init__(connection, path)
        self._connection_service = None

    def add(
        self,
        connection,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('connection', connection, types.StorageConnection),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(connection, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of connections to return. If not specified all the connections are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def connection_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return StorageDomainServerConnectionService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.connection_service(path)
        return self.connection_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'StorageDomainServerConnectionsService:%s' % self._path


class StorageDomainTemplateService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageDomainTemplateService, self).__init__(connection, path)
        self._disks_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def import_(
        self,
        async=None,
        clone=None,
        cluster=None,
        exclusive=None,
        storage_domain=None,
        template=None,
        vm=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Action to import a template from an export storage domain.
        For example, to import the template `456` from the storage domain `123` send the following request:
        [source]
        ----
        POST /ovirt-engine/api/storagedomains/123/templates/456/import
        ----
        With the following request body:
        [source, xml]
        ----
        <action>
          <storage_domain>
            <name>myexport</name>
          </storage_domain>
          <cluster>
            <name>mycluster</name>
          </cluster>
        </action>
        ----


        This method supports the following parameters:

        `clone`:: Use the optional `clone` parameter to generate new UUIDs for the imported template and its entities.
        The user might want to import a template with the `clone` parameter set to `false` when importing a template
        from an export domain, with templates that was exported by a different {product-name} environment.

        `async`:: Indicates if the import should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('clone', clone, bool),
            ('cluster', cluster, types.Cluster),
            ('exclusive', exclusive, bool),
            ('storage_domain', storage_domain, types.StorageDomain),
            ('template', template, types.Template),
            ('vm', vm, types.Vm),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            clone=clone,
            cluster=cluster,
            exclusive=exclusive,
            storage_domain=storage_domain,
            template=template,
            vm=vm,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'import', None, headers, query, wait)

    def register(
        self,
        allow_partial_import=None,
        async=None,
        clone=None,
        cluster=None,
        exclusive=None,
        template=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `allow_partial_import`:: Indicates whether a template is allowed to be registered with only some of its disks.
        If this flag is `true`, the engine will not fail in the validation process if an image is not found, but
        instead it will allow the template to be registered without the missing disks. This is mainly used during
        registration of a template when some of the storage domains are not available. The default value is `false`.

        `async`:: Indicates if the registration should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('allow_partial_import', allow_partial_import, bool),
            ('async', async, bool),
            ('clone', clone, bool),
            ('cluster', cluster, types.Cluster),
            ('exclusive', exclusive, bool),
            ('template', template, types.Template),
        ])

        # Populate the action:
        action = types.Action(
            allow_partial_import=allow_partial_import,
            async=async,
            clone=clone,
            cluster=cluster,
            exclusive=exclusive,
            template=template,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'register', None, headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def disks_service(self):
        """
        """
        return StorageDomainContentDisksService(self._connection, '%s/disks' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'disks':
            return self.disks_service()
        if path.startswith('disks/'):
            return self.disks_service().service(path[6:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StorageDomainTemplateService:%s' % self._path


class StorageDomainTemplatesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageDomainTemplatesService, self).__init__(connection, path)
        self._template_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of templates to return. If not specified all the templates are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def template_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return StorageDomainTemplateService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.template_service(path)
        return self.template_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'StorageDomainTemplatesService:%s' % self._path


class StorageDomainVmService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageDomainVmService, self).__init__(connection, path)
        self._disk_attachments_service = None
        self._disks_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def import_(
        self,
        async=None,
        clone=None,
        cluster=None,
        collapse_snapshots=None,
        storage_domain=None,
        vm=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Imports a virtual machine from an export storage domain.
        For example, send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/storagedomains/123/vms/456/import
        ----
        With a request body like this:
        [source,xml]
        ----
        <action>
          <storage_domain>
            <name>mydata</name>
          </storage_domain>
          <cluster>
            <name>mycluster</name>
          </cluster>
        </action>
        ----
        To import a virtual machine as a new entity add the `clone` parameter:
        [source,xml]
        ----
        <action>
          <storage_domain>
            <name>mydata</name>
          </storage_domain>
          <cluster>
            <name>mycluster</name>
          </cluster>
          <clone>true</clone>
          <vm>
            <name>myvm</name>
          </vm>
        </action>
        ----
        Include an optional `disks` parameter to choose which disks to import. For example, to import the disks
        of the template that have the identifiers `123` and `456` send the following request body:
        [source,xml]
        ----
        <action>
          <cluster>
            <name>mycluster</name>
          </cluster>
          <vm>
            <name>myvm</name>
          </vm>
          <disks>
            <disk id="123"/>
            <disk id="456"/>
          </disks>
        </action>
        ----


        This method supports the following parameters:

        `clone`:: Indicates if the identifiers of the imported virtual machine
        should be regenerated.
        By default when a virtual machine is imported the identifiers
        are preserved. This means that the same virtual machine can't
        be imported multiple times, as that identifiers needs to be
        unique. To allow importing the same machine multiple times set
        this parameter to `true`, as the default is `false`.

        `collapse_snapshots`:: Indicates of the snapshots of the virtual machine that is imported
        should be collapsed, so that the result will be a virtual machine
        without snapshots.
        This parameter is optional, and if it isn't explicitly specified the
        default value is `false`.

        `async`:: Indicates if the import should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('clone', clone, bool),
            ('cluster', cluster, types.Cluster),
            ('collapse_snapshots', collapse_snapshots, bool),
            ('storage_domain', storage_domain, types.StorageDomain),
            ('vm', vm, types.Vm),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            clone=clone,
            cluster=cluster,
            collapse_snapshots=collapse_snapshots,
            storage_domain=storage_domain,
            vm=vm,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'import', None, headers, query, wait)

    def register(
        self,
        allow_partial_import=None,
        async=None,
        clone=None,
        cluster=None,
        reassign_bad_macs=None,
        vm=None,
        vnic_profile_mappings=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `allow_partial_import`:: Indicates whether a virtual machine is allowed to be registered with only some of its disks.
        If this flag is `true`, the engine will not fail in the validation process if an image is not found, but
        instead it will allow the virtual machine to be registered without the missing disks. This is mainly used
        during registration of a virtual machine when some of the storage domains are not available. The default
        value is `false`.

        `vnic_profile_mappings`:: Mapping rules for virtual NIC profiles that will be applied during the import process.

        `reassign_bad_macs`:: Indicates if the problematic MAC addresses should be re-assigned during the import process by the engine.
        A MAC address would be considered as a problematic one if one of the following is true:
        - It conflicts with a MAC address that is already allocated to a virtual machine in the target environment.
        - It's out of the range of the target MAC address pool.

        `async`:: Indicates if the registration should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('allow_partial_import', allow_partial_import, bool),
            ('async', async, bool),
            ('clone', clone, bool),
            ('cluster', cluster, types.Cluster),
            ('reassign_bad_macs', reassign_bad_macs, bool),
            ('vm', vm, types.Vm),
            ('vnic_profile_mappings', vnic_profile_mappings, list),
        ])

        # Populate the action:
        action = types.Action(
            allow_partial_import=allow_partial_import,
            async=async,
            clone=clone,
            cluster=cluster,
            reassign_bad_macs=reassign_bad_macs,
            vm=vm,
            vnic_profile_mappings=vnic_profile_mappings,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'register', None, headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Deletes a virtual machine from an export storage domain.
        For example, to delete the virtual machine `456` from the storage domain `123`, send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/storagedomains/123/vms/456
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def disk_attachments_service(self):
        """
        Returns a reference to the service that manages the disk attachments of the virtual machine.

        """
        return StorageDomainVmDiskAttachmentsService(self._connection, '%s/diskattachments' % self._path)

    def disks_service(self):
        """
        """
        return StorageDomainContentDisksService(self._connection, '%s/disks' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'diskattachments':
            return self.disk_attachments_service()
        if path.startswith('diskattachments/'):
            return self.disk_attachments_service().service(path[16:])
        if path == 'disks':
            return self.disks_service()
        if path.startswith('disks/'):
            return self.disks_service().service(path[6:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StorageDomainVmService:%s' % self._path


class StorageDomainVmDiskAttachmentService(Service):
    """
    Returns the details of the disks attached to a virtual machine in the export domain.

    """

    def __init__(self, connection, path):
        super(StorageDomainVmDiskAttachmentService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns the details of the attachment with all its properties and a link to the disk.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StorageDomainVmDiskAttachmentService:%s' % self._path


class StorageDomainVmDiskAttachmentsService(Service):
    """
    Returns the details of a disk attached to a virtual machine in the export domain.

    """

    def __init__(self, connection, path):
        super(StorageDomainVmDiskAttachmentsService, self).__init__(connection, path)
        self._attachment_service = None

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List the disks that are attached to the virtual machine.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def attachment_service(self, id):
        """
        Reference to the service that manages a specific attachment.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return StorageDomainVmDiskAttachmentService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.attachment_service(path)
        return self.attachment_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'StorageDomainVmDiskAttachmentsService:%s' % self._path


class StorageDomainVmsService(Service):
    """
    Lists the virtual machines of an export storage domain.
    For example, to retrieve the virtual machines that are available in the storage domain with identifier `123` send the
    following request:
    [source]
    ----
    GET /ovirt-engine/api/storagedomains/123/vms
    ----
    This will return the following response body:
    [source,xml]
    ----
    <vms>
      <vm id="456" href="/api/storagedomains/123/vms/456">
        <name>vm1</name>
        ...
        <storage_domain id="123" href="/api/storagedomains/123"/>
        <actions>
          <link rel="import" href="/api/storagedomains/123/vms/456/import"/>
        </actions>
      </vm>
    </vms>
    ----
    Virtual machines and templates in these collections have a similar representation to their counterparts in the
    top-level <<types/vm, Vm>> and <<types/template, Template>> collections, except they also contain a
    <<types/storage_domain, StorageDomain>> reference and an <<services/storage_domain_vm/methods/import, import>>
    action.

    """

    def __init__(self, connection, path):
        super(StorageDomainVmsService, self).__init__(connection, path)
        self._vm_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of virtual machines to return. If not specified all the virtual machines are
        returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def vm_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return StorageDomainVmService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.vm_service(path)
        return self.vm_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'StorageDomainVmsService:%s' % self._path


class StorageDomainsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageDomainsService, self).__init__(connection, path)
        self._storage_domain_service = None

    def add(
        self,
        storage_domain,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adds a new storage domain.
        Creation of a new <<types/storage_domain,StorageDomain>> requires the `name`, `type`, `host` and `storage`
        attributes. Identify the `host` attribute with the `id` or `name` attributes. In oVirt 3.6 and later you can
        enable the wipe after delete option by default on the storage domain. To configure this, specify
        `wipe_after_delete` in the POST request. This option can be edited after the domain is created, but doing so will
        not change the wipe after delete property of disks that already exist.
        To add a new storage domain with specified `name`, `type`, `storage.type`, `storage.address` and `storage.path`
        and by using a host with an id `123`, send a request as follows:
        [source]
        ----
        POST /ovirt-engine/api/storagedomains
        ----
        With a request body as follows:
        [source,xml]
        ----
        <storage_domain>
          <name>mydata</name>
          <type>data</type>
          <storage>
            <type>nfs</type>
            <address>mynfs.example.com</address>
            <path>/exports/mydata</path>
          </storage>
          <host>
            <name>myhost</name>
          </host>
        </storage_domain>
        ----
        To create a new NFS ISO storage domain send a request like this:
        [source,xml]
        ----
        <storage_domain>
          <name>myisos</name>
          <type>iso</type>
          <storage>
            <type>nfs</type>
            <address>mynfs.example.com</address>
            <path>/export/myisos</path>
          </storage>
          <host>
            <name>myhost</name>
          </host>
        </storage_domain>
        ----
        To create a new iSCSI storage domain send a request like this:
        [source,xml]
        ----
        <storage_domain>
          <name>myiscsi</name>
          <type>data</type>
          <storage>
            <type>iscsi</type>
            <logical_units>
              <logical_unit id="3600144f09dbd050000004eedbd340001"/>
              <logical_unit id="3600144f09dbd050000004eedbd340002"/>
            </logical_units>
          </storage>
          <host>
            <name>myhost</name>
          </host>
        </storage_domain>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(storage_domain, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        filter=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of storage domains to return. If not specified all the storage domains are returned.

        `search`:: A query string used to restrict the returned storage domains.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('filter', filter, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def storage_domain_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return StorageDomainService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.storage_domain_service(path)
        return self.storage_domain_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'StorageDomainsService:%s' % self._path


class StorageServerConnectionService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageServerConnectionService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        host=None,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a storage connection.
        A storage connection can only be deleted if neither storage domain nor LUN disks reference it. The host name or
        id is optional; providing it disconnects (unmounts) the connection from that host.


        This method supports the following parameters:

        `host`:: The name or identifier of the host from which the connection would be unmounted (disconnected). If not
        provided, no host will be disconnected.
        For example, to use the host with identifier `456` to delete the storage connection with identifier `123`
        send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/storageconnections/123?host=456
        ----

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('host', host, str),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if host is not None:
            query['host'] = host
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        connection,
        async=None,
        force=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the storage connection.
        For example, to change the address of the storage server send a request like this:
        [source,xml]
        ----
        PUT /ovirt-engine/api/storageconnections/123
        ----
        With a request body like this:
        [source,xml]
        ----
        <storage_connection>
          <address>mynewnfs.example.com</address>
          <host>
            <name>myhost</name>
          </host>
        </storage_connection>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('connection', connection, types.StorageConnection),
            ('async', async, bool),
            ('force', force, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async
        if force is not None:
            force = Writer.render_boolean(force)
            query['force'] = force

        # Send the request and wait for the response:
        return self._internal_update(connection, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StorageServerConnectionService:%s' % self._path


class StorageServerConnectionExtensionService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageServerConnectionExtensionService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        extension,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Update a storage server connection extension for the given host.
        To update the storage connection `456` of host `123` send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/hosts/123/storageconnectionextensions/456
        ----
        With a request body like this:
        [source,xml]
        ----
        <storage_connection_extension>
          <target>iqn.2016-01.com.example:mytarget</target>
          <username>myuser</username>
          <password>mypassword</password>
        </storage_connection_extension>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('extension', extension, types.StorageConnectionExtension),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(extension, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'StorageServerConnectionExtensionService:%s' % self._path


class StorageServerConnectionExtensionsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageServerConnectionExtensionsService, self).__init__(connection, path)
        self._storage_connection_extension_service = None

    def add(
        self,
        extension,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new storage server connection extension for the given host.
        The extension lets the user define credentials for an iSCSI target for a specific host. For example to use
        `myuser` and `mypassword` as the credentials when connecting to the iSCSI target from host `123` send a request
        like this:
        [source]
        ----
        POST /ovirt-engine/api/hosts/123/storageconnectionextensions
        ----
        With a request body like this:
        [source,xml]
        ----
        <storage_connection_extension>
          <target>iqn.2016-01.com.example:mytarget</target>
          <username>myuser</username>
          <password>mypassword</password>
        </storage_connection_extension>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('extension', extension, types.StorageConnectionExtension),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(extension, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of extensions to return. If not specified all the extensions are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def storage_connection_extension_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return StorageServerConnectionExtensionService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.storage_connection_extension_service(path)
        return self.storage_connection_extension_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'StorageServerConnectionExtensionsService:%s' % self._path


class StorageServerConnectionsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(StorageServerConnectionsService, self).__init__(connection, path)
        self._storage_connection_service = None

    def add(
        self,
        connection,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new storage connection.
        For example, to create a new storage connection for the NFS server `mynfs.example.com` and NFS share
        `/export/mydata` send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/storageconnections
        ----
        With a request body like this:
        [source,xml]
        ----
        <storage_connection>
          <type>nfs</type>
          <address>mynfs.example.com</address>
          <path>/export/mydata</path>
          <host>
            <name>myhost</name>
          </host>
        </storage_connection>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('connection', connection, types.StorageConnection),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(connection, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of connections to return. If not specified all the connections are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def storage_connection_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return StorageServerConnectionService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.storage_connection_service(path)
        return self.storage_connection_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'StorageServerConnectionsService:%s' % self._path


class SystemService(Service):
    """
    """

    def __init__(self, connection, path):
        super(SystemService, self).__init__(connection, path)
        self._affinity_labels_service = None
        self._bookmarks_service = None
        self._cluster_levels_service = None
        self._clusters_service = None
        self._cpu_profiles_service = None
        self._data_centers_service = None
        self._disk_profiles_service = None
        self._disks_service = None
        self._domains_service = None
        self._events_service = None
        self._external_host_providers_service = None
        self._external_vm_imports_service = None
        self._groups_service = None
        self._hosts_service = None
        self._icons_service = None
        self._image_transfers_service = None
        self._instance_types_service = None
        self._jobs_service = None
        self._katello_errata_service = None
        self._mac_pools_service = None
        self._network_filters_service = None
        self._networks_service = None
        self._openstack_image_providers_service = None
        self._openstack_network_providers_service = None
        self._openstack_volume_providers_service = None
        self._operating_systems_service = None
        self._permissions_service = None
        self._roles_service = None
        self._scheduling_policies_service = None
        self._scheduling_policy_units_service = None
        self._storage_connections_service = None
        self._storage_domains_service = None
        self._tags_service = None
        self._templates_service = None
        self._users_service = None
        self._vm_pools_service = None
        self._vms_service = None
        self._vnic_profiles_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns basic information describing the API, like the product name, the version number and a summary of the
        number of relevant objects.
        [source]
        ----
        GET /ovirt-engine/api
        ----
        We get following response:
        [source,xml]
        ----
        <api>
          <link rel="capabilities" href="/api/capabilities"/>
          <link rel="clusters" href="/api/clusters"/>
          <link rel="clusters/search" href="/api/clusters?search={query}"/>
          <link rel="datacenters" href="/api/datacenters"/>
          <link rel="datacenters/search" href="/api/datacenters?search={query}"/>
          <link rel="events" href="/api/events"/>
          <link rel="events/search" href="/api/events?search={query}"/>
          <link rel="hosts" href="/api/hosts"/>
          <link rel="hosts/search" href="/api/hosts?search={query}"/>
          <link rel="networks" href="/api/networks"/>
          <link rel="roles" href="/api/roles"/>
          <link rel="storagedomains" href="/api/storagedomains"/>
          <link rel="storagedomains/search" href="/api/storagedomains?search={query}"/>
          <link rel="tags" href="/api/tags"/>
          <link rel="templates" href="/api/templates"/>
          <link rel="templates/search" href="/api/templates?search={query}"/>
          <link rel="users" href="/api/users"/>
          <link rel="groups" href="/api/groups"/>
          <link rel="domains" href="/api/domains"/>
          <link rel="vmpools" href="/api/vmpools"/>
          <link rel="vmpools/search" href="/api/vmpools?search={query}"/>
          <link rel="vms" href="/api/vms"/>
          <link rel="vms/search" href="/api/vms?search={query}"/>
          <product_info>
            <name>oVirt Engine</name>
            <vendor>ovirt.org</vendor>
            <version>
              <build>4</build>
              <full_version>4.0.4</full_version>
              <major>4</major>
              <minor>0</minor>
              <revision>0</revision>
            </version>
          </product_info>
          <special_objects>
            <blank_template href="/ovirt-engine/api/templates/00000000-0000-0000-0000-000000000000" id="00000000-0000-0000-0000-000000000000"/>
            <root_tag href="/ovirt-engine/api/tags/00000000-0000-0000-0000-000000000000" id="00000000-0000-0000-0000-000000000000"/>
          </special_objects>
          <summary>
            <hosts>
              <active>0</active>
              <total>0</total>
            </hosts>
            <storage_domains>
              <active>0</active>
              <total>1</total>
            </storage_domains>
            <users>
              <active>1</active>
              <total>1</total>
            </users>
            <vms>
              <active>0</active>
              <total>0</total>
            </vms>
          </summary>
          <time>2016-09-14T12:00:48.132+02:00</time>
        </api>
        ----
        The entry point provides a user with links to the collections in a
        virtualization environment. The `rel` attribute of each collection link
        provides a reference point for each link.
        The entry point also contains other data such as `product_info`,
        `special_objects` and `summary`.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def reload_configurations(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the reload should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'reloadconfigurations', None, headers, query, wait)

    def affinity_labels_service(self):
        """
        List all known affinity labels.

        """
        return AffinityLabelsService(self._connection, '%s/affinitylabels' % self._path)

    def bookmarks_service(self):
        """
        """
        return BookmarksService(self._connection, '%s/bookmarks' % self._path)

    def cluster_levels_service(self):
        """
        Reference to the service that provides information about the cluster levels supported by the system.

        """
        return ClusterLevelsService(self._connection, '%s/clusterlevels' % self._path)

    def clusters_service(self):
        """
        """
        return ClustersService(self._connection, '%s/clusters' % self._path)

    def cpu_profiles_service(self):
        """
        """
        return CpuProfilesService(self._connection, '%s/cpuprofiles' % self._path)

    def data_centers_service(self):
        """
        """
        return DataCentersService(self._connection, '%s/datacenters' % self._path)

    def disk_profiles_service(self):
        """
        """
        return DiskProfilesService(self._connection, '%s/diskprofiles' % self._path)

    def disks_service(self):
        """
        """
        return DisksService(self._connection, '%s/disks' % self._path)

    def domains_service(self):
        """
        """
        return DomainsService(self._connection, '%s/domains' % self._path)

    def events_service(self):
        """
        """
        return EventsService(self._connection, '%s/events' % self._path)

    def external_host_providers_service(self):
        """
        """
        return ExternalHostProvidersService(self._connection, '%s/externalhostproviders' % self._path)

    def external_vm_imports_service(self):
        """
        Reference to service facilitating import of external virtual machines.

        """
        return ExternalVmImportsService(self._connection, '%s/externalvmimports' % self._path)

    def groups_service(self):
        """
        """
        return GroupsService(self._connection, '%s/groups' % self._path)

    def hosts_service(self):
        """
        """
        return HostsService(self._connection, '%s/hosts' % self._path)

    def icons_service(self):
        """
        """
        return IconsService(self._connection, '%s/icons' % self._path)

    def image_transfers_service(self):
        """
        List of all image transfers being performed for image I/O in oVirt.

        """
        return ImageTransfersService(self._connection, '%s/imagetransfers' % self._path)

    def instance_types_service(self):
        """
        """
        return InstanceTypesService(self._connection, '%s/instancetypes' % self._path)

    def jobs_service(self):
        """
        List all the jobs monitored by the engine.

        """
        return JobsService(self._connection, '%s/jobs' % self._path)

    def katello_errata_service(self):
        """
        List the available Katello errata assigned to the engine.

        """
        return EngineKatelloErrataService(self._connection, '%s/katelloerrata' % self._path)

    def mac_pools_service(self):
        """
        """
        return MacPoolsService(self._connection, '%s/macpools' % self._path)

    def network_filters_service(self):
        """
        Network filters will enhance the admin ability to manage the network packets traffic from/to the participated
        VMs.

        """
        return NetworkFiltersService(self._connection, '%s/networkfilters' % self._path)

    def networks_service(self):
        """
        """
        return NetworksService(self._connection, '%s/networks' % self._path)

    def openstack_image_providers_service(self):
        """
        """
        return OpenstackImageProvidersService(self._connection, '%s/openstackimageproviders' % self._path)

    def openstack_network_providers_service(self):
        """
        """
        return OpenstackNetworkProvidersService(self._connection, '%s/openstacknetworkproviders' % self._path)

    def openstack_volume_providers_service(self):
        """
        """
        return OpenstackVolumeProvidersService(self._connection, '%s/openstackvolumeproviders' % self._path)

    def operating_systems_service(self):
        """
        """
        return OperatingSystemsService(self._connection, '%s/operatingsystems' % self._path)

    def permissions_service(self):
        """
        """
        return SystemPermissionsService(self._connection, '%s/permissions' % self._path)

    def roles_service(self):
        """
        """
        return RolesService(self._connection, '%s/roles' % self._path)

    def scheduling_policies_service(self):
        """
        """
        return SchedulingPoliciesService(self._connection, '%s/schedulingpolicies' % self._path)

    def scheduling_policy_units_service(self):
        """
        """
        return SchedulingPolicyUnitsService(self._connection, '%s/schedulingpolicyunits' % self._path)

    def storage_connections_service(self):
        """
        """
        return StorageServerConnectionsService(self._connection, '%s/storageconnections' % self._path)

    def storage_domains_service(self):
        """
        """
        return StorageDomainsService(self._connection, '%s/storagedomains' % self._path)

    def tags_service(self):
        """
        """
        return TagsService(self._connection, '%s/tags' % self._path)

    def templates_service(self):
        """
        """
        return TemplatesService(self._connection, '%s/templates' % self._path)

    def users_service(self):
        """
        """
        return UsersService(self._connection, '%s/users' % self._path)

    def vm_pools_service(self):
        """
        """
        return VmPoolsService(self._connection, '%s/vmpools' % self._path)

    def vms_service(self):
        """
        """
        return VmsService(self._connection, '%s/vms' % self._path)

    def vnic_profiles_service(self):
        """
        """
        return VnicProfilesService(self._connection, '%s/vnicprofiles' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'affinitylabels':
            return self.affinity_labels_service()
        if path.startswith('affinitylabels/'):
            return self.affinity_labels_service().service(path[15:])
        if path == 'bookmarks':
            return self.bookmarks_service()
        if path.startswith('bookmarks/'):
            return self.bookmarks_service().service(path[10:])
        if path == 'clusterlevels':
            return self.cluster_levels_service()
        if path.startswith('clusterlevels/'):
            return self.cluster_levels_service().service(path[14:])
        if path == 'clusters':
            return self.clusters_service()
        if path.startswith('clusters/'):
            return self.clusters_service().service(path[9:])
        if path == 'cpuprofiles':
            return self.cpu_profiles_service()
        if path.startswith('cpuprofiles/'):
            return self.cpu_profiles_service().service(path[12:])
        if path == 'datacenters':
            return self.data_centers_service()
        if path.startswith('datacenters/'):
            return self.data_centers_service().service(path[12:])
        if path == 'diskprofiles':
            return self.disk_profiles_service()
        if path.startswith('diskprofiles/'):
            return self.disk_profiles_service().service(path[13:])
        if path == 'disks':
            return self.disks_service()
        if path.startswith('disks/'):
            return self.disks_service().service(path[6:])
        if path == 'domains':
            return self.domains_service()
        if path.startswith('domains/'):
            return self.domains_service().service(path[8:])
        if path == 'events':
            return self.events_service()
        if path.startswith('events/'):
            return self.events_service().service(path[7:])
        if path == 'externalhostproviders':
            return self.external_host_providers_service()
        if path.startswith('externalhostproviders/'):
            return self.external_host_providers_service().service(path[22:])
        if path == 'externalvmimports':
            return self.external_vm_imports_service()
        if path.startswith('externalvmimports/'):
            return self.external_vm_imports_service().service(path[18:])
        if path == 'groups':
            return self.groups_service()
        if path.startswith('groups/'):
            return self.groups_service().service(path[7:])
        if path == 'hosts':
            return self.hosts_service()
        if path.startswith('hosts/'):
            return self.hosts_service().service(path[6:])
        if path == 'icons':
            return self.icons_service()
        if path.startswith('icons/'):
            return self.icons_service().service(path[6:])
        if path == 'imagetransfers':
            return self.image_transfers_service()
        if path.startswith('imagetransfers/'):
            return self.image_transfers_service().service(path[15:])
        if path == 'instancetypes':
            return self.instance_types_service()
        if path.startswith('instancetypes/'):
            return self.instance_types_service().service(path[14:])
        if path == 'jobs':
            return self.jobs_service()
        if path.startswith('jobs/'):
            return self.jobs_service().service(path[5:])
        if path == 'katelloerrata':
            return self.katello_errata_service()
        if path.startswith('katelloerrata/'):
            return self.katello_errata_service().service(path[14:])
        if path == 'macpools':
            return self.mac_pools_service()
        if path.startswith('macpools/'):
            return self.mac_pools_service().service(path[9:])
        if path == 'networkfilters':
            return self.network_filters_service()
        if path.startswith('networkfilters/'):
            return self.network_filters_service().service(path[15:])
        if path == 'networks':
            return self.networks_service()
        if path.startswith('networks/'):
            return self.networks_service().service(path[9:])
        if path == 'openstackimageproviders':
            return self.openstack_image_providers_service()
        if path.startswith('openstackimageproviders/'):
            return self.openstack_image_providers_service().service(path[24:])
        if path == 'openstacknetworkproviders':
            return self.openstack_network_providers_service()
        if path.startswith('openstacknetworkproviders/'):
            return self.openstack_network_providers_service().service(path[26:])
        if path == 'openstackvolumeproviders':
            return self.openstack_volume_providers_service()
        if path.startswith('openstackvolumeproviders/'):
            return self.openstack_volume_providers_service().service(path[25:])
        if path == 'operatingsystems':
            return self.operating_systems_service()
        if path.startswith('operatingsystems/'):
            return self.operating_systems_service().service(path[17:])
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'roles':
            return self.roles_service()
        if path.startswith('roles/'):
            return self.roles_service().service(path[6:])
        if path == 'schedulingpolicies':
            return self.scheduling_policies_service()
        if path.startswith('schedulingpolicies/'):
            return self.scheduling_policies_service().service(path[19:])
        if path == 'schedulingpolicyunits':
            return self.scheduling_policy_units_service()
        if path.startswith('schedulingpolicyunits/'):
            return self.scheduling_policy_units_service().service(path[22:])
        if path == 'storageconnections':
            return self.storage_connections_service()
        if path.startswith('storageconnections/'):
            return self.storage_connections_service().service(path[19:])
        if path == 'storagedomains':
            return self.storage_domains_service()
        if path.startswith('storagedomains/'):
            return self.storage_domains_service().service(path[15:])
        if path == 'tags':
            return self.tags_service()
        if path.startswith('tags/'):
            return self.tags_service().service(path[5:])
        if path == 'templates':
            return self.templates_service()
        if path.startswith('templates/'):
            return self.templates_service().service(path[10:])
        if path == 'users':
            return self.users_service()
        if path.startswith('users/'):
            return self.users_service().service(path[6:])
        if path == 'vmpools':
            return self.vm_pools_service()
        if path.startswith('vmpools/'):
            return self.vm_pools_service().service(path[8:])
        if path == 'vms':
            return self.vms_service()
        if path.startswith('vms/'):
            return self.vms_service().service(path[4:])
        if path == 'vnicprofiles':
            return self.vnic_profiles_service()
        if path.startswith('vnicprofiles/'):
            return self.vnic_profiles_service().service(path[13:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'SystemService:%s' % self._path


class SystemPermissionsService(AssignedPermissionsService):
    """
    This service doesn't add any new methods, it is just a placeholder for the annotation that specifies the path of the
    resource that manages the permissions assigned to the system object.

    """

    def __init__(self, connection, path):
        super(SystemPermissionsService, self).__init__(connection, path)
        self._permission_service = None

    def add(
        self,
        permission,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Assign a new permission to a user or group for specific entity.
        For example, to assign the `UserVmManager` role to the virtual machine with id `123` to the user with id `456`
        send a request like this:
        ....
        POST /ovirt-engine/api/vms/123/permissions
        ....
        With a request body like this:
        [source,xml]
        ----
        <permission>
          <role>
            <name>UserVmManager</name>
          </role>
          <user id="456"/>
        </permission>
        ----
        To assign the `SuperUser` role to the system to the user with id `456` send a request like this:
        ....
        POST /ovirt-engine/api/permissions
        ....
        With a request body like this:
        [source,xml]
        ----
        <permission>
          <role>
            <name>SuperUser</name>
          </role>
          <user id="456"/>
        </permission>
        ----
        If you want to assign permission to the group instead of the user please replace the `user` element with the
        `group` element with proper `id` of the group. For example to assign the `UserRole` role to the cluster with
        id `123` to the group with id `789` send a request like this:
        ....
        POST /ovirt-engine/api/clusters/123/permissions
        ....
        With a request body like this:
        [source,xml]
        ----
        <permission>
          <role>
            <name>UserRole</name>
          </role>
          <group id="789"/>
        </permission>
        ----


        This method supports the following parameters:

        `permission`:: The permission.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('permission', permission, types.Permission),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(permission, headers, query, wait)

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List all the permissions of the specific entity.
        For example to list all the permissions of the cluster with id `123` send a request like this:
        ....
        GET /ovirt-engine/api/clusters/123/permissions
        ....
        [source,xml]
        ----
        <permissions>
          <permission id="456">
            <cluster id="123"/>
            <role id="789"/>
            <user id="451"/>
          </permission>
          <permission id="654">
            <cluster id="123"/>
            <role id="789"/>
            <group id="127"/>
          </permission>
        </permissions>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def permission_service(self, id):
        """
        Sub-resource locator method, returns individual permission resource on which the remainder of the URI is
        dispatched.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return PermissionService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.permission_service(path)
        return self.permission_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'SystemPermissionsService:%s' % self._path


class TagService(Service):
    """
    A service to manage a specific tag in the system.

    """

    def __init__(self, connection, path):
        super(TagService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets the information about the tag.
        For example to retrieve the information about the tag with the id `123` send a request like this:
        ....
        GET /ovirt-engine/api/tags/123
        ....
        [source,xml]
        ----
        <tag href="/ovirt-engine/api/tags/123" id="123">
          <name>root</name>
          <description>root</description>
        </tag>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the tag from the system.
        For example to remove the tag with id `123` send a request like this:
        ....
        DELETE /ovirt-engine/api/tags/123
        ....


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        tag,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the tag entity.
        For example to update parent tag to tag with id `456` of the tag with id `123` send a request like this:
        ....
        PUT /ovirt-engine/api/tags/123
        ....
        With request body like:
        [source,xml]
        ----
        <tag>
          <parent id="456"/>
        </tag>
        ----
        You may also specify a tag name instead of id. For example to update parent tag to tag with name `mytag`
        of the tag with id `123` send a request like this:
        [source,xml]
        ----
        <tag>
          <parent>
            <name>mytag</name>
          </parent>
        </tag>
        ----


        This method supports the following parameters:

        `tag`:: The updated tag.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('tag', tag, types.Tag),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(tag, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'TagService:%s' % self._path


class TagsService(Service):
    """
    Represents a service to manage collection of the tags in the system.

    """

    def __init__(self, connection, path):
        super(TagsService, self).__init__(connection, path)
        self._tag_service = None

    def add(
        self,
        tag,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add a new tag to the system.
        For example, to add new tag with name `mytag` to the system send a request like this:
        ....
        POST /ovirt-engine/api/tags
        ....
        With a request body like this:
        [source,xml]
        ----
        <tag>
          <name>mytag</name>
        </tag>
        ----
        NOTE: The root tag is a special pseudo-tag assumed as the default parent tag if no parent tag is specified.
        The root tag cannot be deleted nor assigned a parent tag.
        To create new tag with specific parent tag send a request body like this:
        [source,xml]
        ----
        <tag>
          <name>mytag</name>
          <parent>
            <name>myparenttag</name>
          </parent>
        </tag>
        ----


        This method supports the following parameters:

        `tag`:: The added tag.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('tag', tag, types.Tag),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(tag, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List the tags in the system.
        For example to list the full hierarchy of the tags in the system send a request like this:
        ....
        GET /ovirt-engine/api/tags
        ....
        [source,xml]
        ----
        <tags>
          <tag href="/ovirt-engine/api/tags/222" id="222">
            <name>root2</name>
            <description>root2</description>
            <parent href="/ovirt-engine/api/tags/111" id="111"/>
          </tag>
          <tag href="/ovirt-engine/api/tags/333" id="333">
            <name>root3</name>
            <description>root3</description>
            <parent href="/ovirt-engine/api/tags/222" id="222"/>
          </tag>
          <tag href="/ovirt-engine/api/tags/111" id="111">
            <name>root</name>
            <description>root</description>
          </tag>
        </tags>
        ----
        In the previous XML output you can see the following hierarchy of the tags:
        ....
        root:        (id: 111)
          - root2    (id: 222)
            - root3  (id: 333)
        ....


        This method supports the following parameters:

        `max`:: Sets the maximum number of tags to return. If not specified all the tags are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def tag_service(self, id):
        """
        Reference to the service that manages a specific tag.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return TagService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.tag_service(path)
        return self.tag_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'TagsService:%s' % self._path


class TemplateService(Service):
    """
    Manages the virtual machine template and template versions.

    """

    def __init__(self, connection, path):
        super(TemplateService, self).__init__(connection, path)
        self._cdroms_service = None
        self._disk_attachments_service = None
        self._graphics_consoles_service = None
        self._nics_service = None
        self._permissions_service = None
        self._tags_service = None
        self._watchdogs_service = None

    def export(
        self,
        exclusive=None,
        storage_domain=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Exports a template to the data center export domain.
        For example, the operation can be facilitated using the following request:
        [source]
        ----
        POST /ovirt-engine/api/templates/123/export
        ----
        With a request body like this:
        [source,xml]
        ----
        <action>
          <storage_domain id="456"/>
          <exclusive>true<exclusive/>
        </action>
        ----


        This method supports the following parameters:

        `exclusive`:: Indicates if the existing templates with the same name should be overwritten.
        The export action reports a failed action if a template of the same name exists in the destination domain.
        Set this parameter to `true` to change this behavior and overwrite any existing template.

        `storage_domain`:: Specifies the destination export storage domain.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('exclusive', exclusive, bool),
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Populate the action:
        action = types.Action(
            exclusive=exclusive,
            storage_domain=storage_domain,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'export', None, headers, query, wait)

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns the information about this template or template version.


        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a virtual machine template.
        [source]
        ----
        DELETE /ovirt-engine/api/templates/123
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def seal(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Seal the template.
        Sealing erases all host-specific configuration from the filesystem:
        SSH keys, UDEV rules, MAC addresses, system ID, hostname etc.,
        thus making easy to use the template to create multiple virtual
        machines without manual intervention.
        Currently sealing is supported only for Linux OS.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'seal', None, headers, query, wait)

    def update(
        self,
        template,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the template.
        The `name`, `description`, `type`, `memory`, `cpu`, `topology`, `os`, `high_availability`, `display`,
        `stateless`, `usb` and `timezone` elements can be updated after a template has been created.
        For example, to update a template to so that it has 1 GiB of memory send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/templates/123
        ----
        With the following request body:
        [source,xml]
        ----
        <template>
          <memory>1073741824</memory>
        </template>
        ----
        The `version_name` name attribute is the only one that can be updated within the `version` attribute used for
        template versions:
        [source,xml]
        ----
        <template>
          <version>
            <version_name>mytemplate_2</version_name>
          </version>
        </template>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('template', template, types.Template),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(template, headers, query, wait)

    def cdroms_service(self):
        """
        Returns a reference to the service that manages the CDROMs that are associated with the template.

        """
        return TemplateCdromsService(self._connection, '%s/cdroms' % self._path)

    def disk_attachments_service(self):
        """
        Reference to the service that manages a specific
        disk attachment of the template.

        """
        return TemplateDiskAttachmentsService(self._connection, '%s/diskattachments' % self._path)

    def graphics_consoles_service(self):
        """
        Returns a reference to the service that manages the graphical consoles that are associated with the template.

        """
        return TemplateGraphicsConsolesService(self._connection, '%s/graphicsconsoles' % self._path)

    def nics_service(self):
        """
        Returns a reference to the service that manages the NICs that are associated with the template.

        """
        return TemplateNicsService(self._connection, '%s/nics' % self._path)

    def permissions_service(self):
        """
        Returns a reference to the service that manages the permissions that are associated with the template.

        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def tags_service(self):
        """
        Returns a reference to the service that manages the tags that are associated with the template.

        """
        return AssignedTagsService(self._connection, '%s/tags' % self._path)

    def watchdogs_service(self):
        """
        Returns a reference to the service that manages the _watchdogs_ that are associated with the template.

        """
        return TemplateWatchdogsService(self._connection, '%s/watchdogs' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'cdroms':
            return self.cdroms_service()
        if path.startswith('cdroms/'):
            return self.cdroms_service().service(path[7:])
        if path == 'diskattachments':
            return self.disk_attachments_service()
        if path.startswith('diskattachments/'):
            return self.disk_attachments_service().service(path[16:])
        if path == 'graphicsconsoles':
            return self.graphics_consoles_service()
        if path.startswith('graphicsconsoles/'):
            return self.graphics_consoles_service().service(path[17:])
        if path == 'nics':
            return self.nics_service()
        if path.startswith('nics/'):
            return self.nics_service().service(path[5:])
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'tags':
            return self.tags_service()
        if path.startswith('tags/'):
            return self.tags_service().service(path[5:])
        if path == 'watchdogs':
            return self.watchdogs_service()
        if path.startswith('watchdogs/'):
            return self.watchdogs_service().service(path[10:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'TemplateService:%s' % self._path


class TemplateCdromService(Service):
    """
    A service managing a CD-ROM device on templates.

    """

    def __init__(self, connection, path):
        super(TemplateCdromService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns the information about this CD-ROM device.
        For example, to get information about the CD-ROM device of template `123` send a request like:
        [source]
        ----
        GET /ovirt-engine/api/templates/123/cdroms/
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'TemplateCdromService:%s' % self._path


class TemplateCdromsService(Service):
    """
    Lists the CD-ROM devices of a template.

    """

    def __init__(self, connection, path):
        super(TemplateCdromsService, self).__init__(connection, path)
        self._cdrom_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of CD-ROMs to return. If not specified all the CD-ROMs are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def cdrom_service(self, id):
        """
        Returns a reference to the service that manages a specific CD-ROM device.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return TemplateCdromService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.cdrom_service(path)
        return self.cdrom_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'TemplateCdromsService:%s' % self._path


class TemplateDiskService(Service):
    """
    """

    def __init__(self, connection, path):
        super(TemplateDiskService, self).__init__(connection, path)

    def copy(
        self,
        async=None,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the copy should be performed asynchronously.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('filter', filter, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            filter=filter,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'copy', None, headers, query, wait)

    def export(
        self,
        async=None,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the export should be performed asynchronously.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('filter', filter, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            filter=filter,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'export', None, headers, query, wait)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'TemplateDiskService:%s' % self._path


class TemplateDiskAttachmentService(Service):
    """
    This service manages the attachment of a disk to a template.

    """

    def __init__(self, connection, path):
        super(TemplateDiskAttachmentService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns the details of the attachment.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        storage_domain=None,
        force=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the disk from the template. The disk will only be removed if there are other existing copies of the
        disk on other storage domains.
        A storage domain has to be specified to determine which of the copies should be removed (template disks can
        have copies on multiple storage domains).
        [source]
        ----
        DELETE /ovirt-engine/api/templates/{template:id}/diskattachments/{attachment:id}?storage_domain=072fbaa1-08f3-4a40-9f34-a5ca22dd1d74
        ----


        This method supports the following parameters:

        `storage_domain`:: Specifies the identifier of the storage domain the image to be removed resides on.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('storage_domain', storage_domain, str),
            ('force', force, bool),
        ])

        # Build the URL:
        query = query or {}
        if storage_domain is not None:
            query['storage_domain'] = storage_domain
        if force is not None:
            force = Writer.render_boolean(force)
            query['force'] = force

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'TemplateDiskAttachmentService:%s' % self._path


class TemplateDiskAttachmentsService(Service):
    """
    This service manages the set of disks attached to a template. Each attached disk is represented by a
    <<types/disk_attachment,DiskAttachment>>.

    """

    def __init__(self, connection, path):
        super(TemplateDiskAttachmentsService, self).__init__(connection, path)
        self._attachment_service = None

    def list(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List the disks that are attached to the template.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def attachment_service(self, id):
        """
        Reference to the service that manages a specific attachment.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return TemplateDiskAttachmentService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.attachment_service(path)
        return self.attachment_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'TemplateDiskAttachmentsService:%s' % self._path


class TemplateDisksService(Service):
    """
    """

    def __init__(self, connection, path):
        super(TemplateDisksService, self).__init__(connection, path)
        self._disk_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def disk_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return TemplateDiskService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.disk_service(path)
        return self.disk_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'TemplateDisksService:%s' % self._path


class TemplateGraphicsConsoleService(Service):
    """
    """

    def __init__(self, connection, path):
        super(TemplateGraphicsConsoleService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets graphics console configuration of the template.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove the graphics console from the template.


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'TemplateGraphicsConsoleService:%s' % self._path


class TemplateGraphicsConsolesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(TemplateGraphicsConsolesService, self).__init__(connection, path)
        self._console_service = None

    def add(
        self,
        console,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add new graphics console to the template.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('console', console, types.GraphicsConsole),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(console, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists all the configured graphics consoles of the template.


        This method supports the following parameters:

        `max`:: Sets the maximum number of consoles to return. If not specified all the consoles are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def console_service(self, id):
        """
        Returns a reference to the service that manages a specific template graphics console.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return TemplateGraphicsConsoleService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.console_service(path)
        return self.console_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'TemplateGraphicsConsolesService:%s' % self._path


class TemplateNicService(Service):
    """
    """

    def __init__(self, connection, path):
        super(TemplateNicService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        nic,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('nic', nic, types.Nic),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(nic, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'TemplateNicService:%s' % self._path


class TemplateNicsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(TemplateNicsService, self).__init__(connection, path)
        self._nic_service = None

    def add(
        self,
        nic,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('nic', nic, types.Nic),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(nic, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of NICs to return. If not specified all the NICs are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def nic_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return TemplateNicService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.nic_service(path)
        return self.nic_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'TemplateNicsService:%s' % self._path


class TemplateWatchdogService(Service):
    """
    """

    def __init__(self, connection, path):
        super(TemplateWatchdogService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        watchdog,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('watchdog', watchdog, types.Watchdog),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(watchdog, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'TemplateWatchdogService:%s' % self._path


class TemplateWatchdogsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(TemplateWatchdogsService, self).__init__(connection, path)
        self._watchdog_service = None

    def add(
        self,
        watchdog,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('watchdog', watchdog, types.Watchdog),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(watchdog, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of watchdogs to return. If not specified all the watchdogs are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def watchdog_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return TemplateWatchdogService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.watchdog_service(path)
        return self.watchdog_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'TemplateWatchdogsService:%s' % self._path


class TemplatesService(Service):
    """
    This service manages the virtual machine templates available in the system.

    """

    def __init__(self, connection, path):
        super(TemplatesService, self).__init__(connection, path)
        self._template_service = None

    def add(
        self,
        template,
        clone_permissions=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new template.
        This requires the `name` and `vm` elements. Identify the virtual machine with the `id` `name` attributes.
        [source]
        ----
        POST /ovirt-engine/api/templates
        ----
        With a request body like this:
        [source,xml]
        ----
        <template>
          <name>mytemplate</name>
          <vm id="123"/>
        </template>
        ----
        The template can be created as a sub version of an existing template.This requires the `name` and `vm` attributes
        for the new template, and the `base_template` and `version_name` attributes for the new template version. The
        `base_template` and `version_name` attributes must be specified within a `version` section enclosed in the
        `template` section. Identify the virtual machine with the `id` or `name` attributes.
        [source,xml]
        ----
        <template>
          <name>mytemplate</name>
          <vm id="123"/>
          <version>
            <base_template id="456"/>
            <version_name>mytemplate_001</version_name>
          </version>
        </template>
        ----


        This method supports the following parameters:

        `template`:: The information about the template or template version.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('template', template, types.Template),
            ('clone_permissions', clone_permissions, bool),
        ])

        # Build the URL:
        query = query or {}
        if clone_permissions is not None:
            clone_permissions = Writer.render_boolean(clone_permissions)
            query['clone_permissions'] = clone_permissions

        # Send the request and wait for the response:
        return self._internal_add(template, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        filter=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns the list of virtual machine templates.
        For example:
        [source]
        ----
        GET /ovirt-engine/api/templates
        ----
        Will return the list of virtual machines and virtual machine templates.


        This method supports the following parameters:

        `max`:: Sets the maximum number of templates to return. If not specified all the templates are returned.

        `search`:: A query string used to restrict the returned templates.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('filter', filter, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def template_service(self, id):
        """
        Returns a reference to the service that manages a specific virtual machine template.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return TemplateService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.template_service(path)
        return self.template_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'TemplatesService:%s' % self._path


class UnmanagedNetworkService(Service):
    """
    """

    def __init__(self, connection, path):
        super(UnmanagedNetworkService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'UnmanagedNetworkService:%s' % self._path


class UnmanagedNetworksService(Service):
    """
    """

    def __init__(self, connection, path):
        super(UnmanagedNetworksService, self).__init__(connection, path)
        self._unmanaged_network_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def unmanaged_network_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return UnmanagedNetworkService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.unmanaged_network_service(path)
        return self.unmanaged_network_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'UnmanagedNetworksService:%s' % self._path


class UserService(Service):
    """
    A service to manage a user in the system.
    Use this service to either get users details or remove users.
    In order to add new users please use
    <<services/users>>.

    """

    def __init__(self, connection, path):
        super(UserService, self).__init__(connection, path)
        self._permissions_service = None
        self._roles_service = None
        self._ssh_public_keys_service = None
        self._tags_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets the system user information.
        Usage:
        ....
        GET /ovirt-engine/api/users/1234
        ....
        Will return the user information:
        [source,xml]
        ----
        <user href="/ovirt-engine/api/users/1234" id="1234">
          <name>admin</name>
          <link href="/ovirt-engine/api/users/1234/sshpublickeys" rel="sshpublickeys"/>
          <link href="/ovirt-engine/api/users/1234/roles" rel="roles"/>
          <link href="/ovirt-engine/api/users/1234/permissions" rel="permissions"/>
          <link href="/ovirt-engine/api/users/1234/tags" rel="tags"/>
          <department></department>
          <domain_entry_id>23456</domain_entry_id>
          <email>user1@domain.com</email>
          <last_name>Lastname</last_name>
          <namespace>*</namespace>
          <principal>user1</principal>
          <user_name>user1@domain-authz</user_name>
          <domain href="/ovirt-engine/api/domains/45678" id="45678">
            <name>domain-authz</name>
          </domain>
        </user>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the system user.
        Usage:
        ....
        DELETE /ovirt-engine/api/users/1234
        ....


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def permissions_service(self):
        """
        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def roles_service(self):
        """
        """
        return AssignedRolesService(self._connection, '%s/roles' % self._path)

    def ssh_public_keys_service(self):
        """
        """
        return SshPublicKeysService(self._connection, '%s/sshpublickeys' % self._path)

    def tags_service(self):
        """
        """
        return AssignedTagsService(self._connection, '%s/tags' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'roles':
            return self.roles_service()
        if path.startswith('roles/'):
            return self.roles_service().service(path[6:])
        if path == 'sshpublickeys':
            return self.ssh_public_keys_service()
        if path.startswith('sshpublickeys/'):
            return self.ssh_public_keys_service().service(path[14:])
        if path == 'tags':
            return self.tags_service()
        if path.startswith('tags/'):
            return self.tags_service().service(path[5:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'UserService:%s' % self._path


class UsersService(Service):
    """
    A service to manage the users in the system.

    """

    def __init__(self, connection, path):
        super(UsersService, self).__init__(connection, path)
        self._user_service = None

    def add(
        self,
        user,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add user from a directory service.
        For example, to add the `myuser` user from the `myextension-authz` authorization provider send a request
        like this:
        [source]
        ----
        POST /ovirt-engine/api/users
        ----
        With a request body like this:
        [source,xml]
        ----
        <user>
          <user_name>myuser@myextension-authz</user_name>
          <domain>
            <name>myextension-authz</name>
          </domain>
        </user>
        ----
        In case you are working with Active Directory you have to pass user principal name (UPN) as `username`, followed
        by authorization provider name. Due to https://bugzilla.redhat.com/1147900[bug 1147900] you need to provide
        also `principal` parameter set to UPN of the user.
        For example, to add the user with UPN `myuser@mysubdomain.mydomain.com` from the `myextension-authz`
        authorization provider send a request body like this:
        [source,xml]
        ----
        <user>
          <principal>myuser@mysubdomain.mydomain.com</principal>
          <user_name>myuser@mysubdomain.mydomain.com@myextension-authz</user_name>
          <domain>
            <name>myextension-authz</name>
          </domain>
        </user>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('user', user, types.User),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(user, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List all the users in the system.
        Usage:
        ....
        GET /ovirt-engine/api/users
        ....
        Will return the list of users:
        [source,xml]
        ----
        <users>
          <user href="/ovirt-engine/api/users/1234" id="1234">
            <name>admin</name>
            <link href="/ovirt-engine/api/users/1234/sshpublickeys" rel="sshpublickeys"/>
            <link href="/ovirt-engine/api/users/1234/roles" rel="roles"/>
            <link href="/ovirt-engine/api/users/1234/permissions" rel="permissions"/>
            <link href="/ovirt-engine/api/users/1234/tags" rel="tags"/>
            <domain_entry_id>23456</domain_entry_id>
            <namespace>*</namespace>
            <principal>user1</principal>
            <user_name>user1@domain-authz</user_name>
            <domain href="/ovirt-engine/api/domains/45678" id="45678">
              <name>domain-authz</name>
            </domain>
          </user>
        </users>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of users to return. If not specified all the users are returned.

        `search`:: A query string used to restrict the returned users.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def user_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return UserService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.user_service(path)
        return self.user_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'UsersService:%s' % self._path


class VirtualFunctionAllowedNetworkService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VirtualFunctionAllowedNetworkService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VirtualFunctionAllowedNetworkService:%s' % self._path


class VirtualFunctionAllowedNetworksService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VirtualFunctionAllowedNetworksService, self).__init__(connection, path)
        self._network_service = None

    def add(
        self,
        network,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('network', network, types.Network),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(network, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of networks to return. If not specified all the networks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def network_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return VirtualFunctionAllowedNetworkService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.network_service(path)
        return self.network_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VirtualFunctionAllowedNetworksService:%s' % self._path


class VmService(MeasurableService):
    """
    """

    def __init__(self, connection, path):
        super(VmService, self).__init__(connection, path)
        self._affinity_labels_service = None
        self._applications_service = None
        self._cdroms_service = None
        self._disk_attachments_service = None
        self._graphics_consoles_service = None
        self._host_devices_service = None
        self._katello_errata_service = None
        self._nics_service = None
        self._numa_nodes_service = None
        self._permissions_service = None
        self._reported_devices_service = None
        self._sessions_service = None
        self._snapshots_service = None
        self._statistics_service = None
        self._tags_service = None
        self._watchdogs_service = None

    def cancel_migration(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation stops any migration of a virtual machine to another physical host.
        [source]
        ----
        POST /ovirt-engine/api/vms/123/cancelmigration
        ----
        The cancel migration action does not take any action specific parameters,
        so the request body should contain an empty `action`:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the migration should cancelled asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'cancelmigration', None, headers, query, wait)

    def clone(
        self,
        async=None,
        vm=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the clone should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('vm', vm, types.Vm),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            vm=vm,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'clone', None, headers, query, wait)

    def commit_snapshot(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the snapshots should be committed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'commitsnapshot', None, headers, query, wait)

    def detach(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Detaches a virtual machine from a pool.
        [source]
        ----
        POST /ovirt-engine/api/vms/123/detach
        ----
        The detach action does not take any action specific parameters, so the request body should contain an
        empty `action`:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the detach should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'detach', None, headers, query, wait)

    def export(
        self,
        async=None,
        discard_snapshots=None,
        exclusive=None,
        storage_domain=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Export a virtual machine to an export domain.
        For example to export virtual machine `123` to the export domain `myexport`, send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/vms/123/export
        ----
        With a request body like this:
        [source,xml]
        ----
        <action>
          <storage_domain>
            <name>myexport</name>
          </storage_domain>
          <exclusive>true</exclusive>
          <discard_snapshots>true</discard_snapshots>
        </action>
        ----


        This method supports the following parameters:

        `discard_snapshots`:: The `discard_snapshots` parameter is to be used when the virtual machine should be exported with all its
        snapshots collapsed.

        `exclusive`:: The `exclusive` parameter is to be used when the virtual machine should be exported even if another copy of
        it already exists in the export domain (override).

        `async`:: Indicates if the export should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('discard_snapshots', discard_snapshots, bool),
            ('exclusive', exclusive, bool),
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            discard_snapshots=discard_snapshots,
            exclusive=exclusive,
            storage_domain=storage_domain,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'export', None, headers, query, wait)

    def freeze_filesystems(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Freeze virtual machine file systems.
        This operation freezes a virtual machine's file systems using the QEMU guest agent when taking a live snapshot of
        a running virtual machine. Normally, this is done automatically by the manager, but this must be executed
        manually with the API for virtual machines using OpenStack Volume (Cinder) disks.
        Example:
        [source]
        ----
        POST /ovirt-engine/api/vms/123/freezefilesystems
        ----
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the freeze should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'freezefilesystems', None, headers, query, wait)

    def get(
        self,
        all_content=None,
        filter=None,
        next_run=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the description of the virtual machine.


        This method supports the following parameters:

        `next_run`:: Indicates if the returned result describes the virtual machine as it is currently running, or if describes
        it with the modifications that have already been performed but that will have effect only when it is
        restarted. By default the values is `false`.
        If the parameter is included in the request, but without a value, it is assumed that the value is `true`, so
        the following request:
        [source]
        ----
        GET /vms/{vm:id};next_run
        ----
        Is equivalent to using the value `true`:
        [source]
        ----
        GET /vms/{vm:id};next_run=true
        ----

        `all_content`:: Indicates if all the attributes of the virtual machine should be included in the response.
        By default the following attributes are excluded:
        - `console`
        - `initialization.configuration.data` - The OVF document describing the virtual machine.
        - `rng_source`
        - `soundcard`
        - `virtio_scsi`
        For example, to retrieve the complete representation of the virtual machine '123' send a request like this:
        ....
        GET /ovirt-engine/api/vms/123?all_content=true
        ....
        NOTE: The reason for not including these attributes is performance: they are seldom used and they require
        additional queries to the database. So try to use the this parameter only when it is really needed.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('all_content', all_content, bool),
            ('filter', filter, bool),
            ('next_run', next_run, bool),
        ])

        # Build the URL:
        query = query or {}
        if all_content is not None:
            all_content = Writer.render_boolean(all_content)
            query['all_content'] = all_content
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if next_run is not None:
            next_run = Writer.render_boolean(next_run)
            query['next_run'] = next_run

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def logon(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Initiates the automatic user logon to access a virtual machine from an external console.
        This action requires the `ovirt-guest-agent-gdm-plugin` and the `ovirt-guest-agent-pam-module` packages to be
        installed and the `ovirt-guest-agent` service to be running on the virtual machine.
        Users require the appropriate user permissions for the virtual machine in order to access the virtual machine
        from an external console.
        This is how an example request would look like:
        [source]
        ----
        POST /ovirt-engine/api/vms/123/logon
        ----
        Request body:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the logon should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'logon', None, headers, query, wait)

    def maintenance(
        self,
        async=None,
        maintenance_enabled=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Sets the global maintenance mode on the hosted engine virtual machine.
        This action has no effect on other virtual machines.
        Example:
        [source]
        ----
        POST /ovirt-engine/api/vms/123/maintenance
        ----
        [source,xml]
        ----
        <action>
          <maintenance_enabled>true<maintenance_enabled/>
        </action>
        ----


        This method supports the following parameters:

        `maintenance_enabled`:: Indicates if global maintenance should be enabled or disabled.

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('maintenance_enabled', maintenance_enabled, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            maintenance_enabled=maintenance_enabled,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'maintenance', None, headers, query, wait)

    def migrate(
        self,
        async=None,
        cluster=None,
        force=None,
        host=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation migrates a virtual machine to another physical host.
        [source]
        ----
        POST /ovirt-engine/api/vms/123/migrate
        ----
        One can specify a specific host to migrate the virtual machine to:
        [source,xml]
        ----
        <action>
          <host id="2ab5e1da-b726-4274-bbf7-0a42b16a0fc3"/>
        </action>
        ----


        This method supports the following parameters:

        `cluster`:: Specifies the cluster the virtual machine should migrate to. This is an optional parameter. By default, the
        virtual machine is migrated to another host within the same cluster.

        `force`:: Specifies the virtual machine should migrate although it might be defined as non migratable. This is an
        optional parameter. By default, it is set to `false`.

        `host`:: Specifies a specific host the virtual machine should migrate to. This is an optional parameters. By default,
        the oVirt Engine automatically selects a default host for migration within the same cluster. If an API user
        requires a specific host, the user can specify the host with either an `id` or `name` parameter.

        `async`:: Indicates if the migration should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('cluster', cluster, types.Cluster),
            ('force', force, bool),
            ('host', host, types.Host),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            cluster=cluster,
            force=force,
            host=host,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'migrate', None, headers, query, wait)

    def preview_snapshot(
        self,
        async=None,
        disks=None,
        restore_memory=None,
        snapshot=None,
        vm=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the preview should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('disks', disks, list),
            ('restore_memory', restore_memory, bool),
            ('snapshot', snapshot, types.Snapshot),
            ('vm', vm, types.Vm),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            disks=disks,
            restore_memory=restore_memory,
            snapshot=snapshot,
            vm=vm,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'previewsnapshot', None, headers, query, wait)

    def reboot(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation sends a reboot request to a virtual machine.
        [source]
        ----
        POST /ovirt-engine/api/vms/123/reboot
        ----
        The reboot action does not take any action specific parameters, so the request body should contain an
        empty `action`:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the reboot should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'reboot', None, headers, query, wait)

    def remove(
        self,
        async=None,
        detach_only=None,
        force=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the virtual machine, including the virtual disks attached to it.
        For example, to remove the virtual machine with identifier `123` send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/vms/123
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `detach_only`:: Indicates if the attached virtual disks should be detached first and preserved instead of being removed.

        `force`:: Indicates if the virtual machine should be forcibly removed.
        Locked virtual machines and virtual machines with locked disk images
        cannot be removed without this flag set to true.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('detach_only', detach_only, bool),
            ('force', force, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async
        if detach_only is not None:
            detach_only = Writer.render_boolean(detach_only)
            query['detach_only'] = detach_only
        if force is not None:
            force = Writer.render_boolean(force)
            query['force'] = force

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def reorder_mac_addresses(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'reordermacaddresses', None, headers, query, wait)

    def shutdown(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation sends a shutdown request to a virtual machine.
        [source]
        ----
        POST /ovirt-engine/api/vms/123/shutdown
        ----
        The shutdown action does not take any action specific parameters,
        so the request body should contain an empty `action`:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the shutdown should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'shutdown', None, headers, query, wait)

    def start(
        self,
        async=None,
        filter=None,
        pause=None,
        use_cloud_init=None,
        use_sysprep=None,
        vm=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Starts the virtual machine.
        If the virtual environment is complete and the virtual machine contains all necessary components to function,
        it can be started.
        This example starts the virtual machine:
        [source]
        ----
        POST /ovirt-engine/api/vms/123/start
        ----
        With a request body:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `pause`:: If set to `true`, start the virtual machine in paused mode. Default is `false`.

        `vm`:: The definition of the virtual machine for this specific run.
        For example:
        [source,xml]
        ----
        <action>
          <vm>
            <os>
              <boot>
                <devices>
                  <device>cdrom</device>
                </devices>
              </boot>
            </os>
          </vm>
        </action>
        ----
        This will set the boot device to the CDROM only for this specific start. After the virtual machine will be
        powered off, this definition will be reverted.

        `use_cloud_init`:: If set to `true`, the initialization type is set to _cloud-init_. The default value is `false`.
        See https://cloudinit.readthedocs.io/en/latest[this] for details.

        `use_sysprep`:: If set to `true`, the initialization type is set to _Sysprep_. The default value is `false`.
        See https://en.wikipedia.org/wiki/Sysprep[this] for details.

        `async`:: Indicates if the action should be performed asynchronously.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('filter', filter, bool),
            ('pause', pause, bool),
            ('use_cloud_init', use_cloud_init, bool),
            ('use_sysprep', use_sysprep, bool),
            ('vm', vm, types.Vm),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            filter=filter,
            pause=pause,
            use_cloud_init=use_cloud_init,
            use_sysprep=use_sysprep,
            vm=vm,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'start', None, headers, query, wait)

    def stop(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation forces a virtual machine to power-off.
        [source]
        ----
        POST /ovirt-engine/api/vms/123/stop
        ----
        The stop action does not take any action specific parameters,
        so the request body should contain an empty `action`:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'stop', None, headers, query, wait)

    def suspend(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation saves the virtual machine state to disk and stops it.
        Start a suspended virtual machine and restore the virtual machine state with the start action.
        [source]
        ----
        POST /ovirt-engine/api/vms/123/suspend
        ----
        The suspend action does not take any action specific parameters,
        so the request body should contain an empty `action`:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'suspend', None, headers, query, wait)

    def thaw_filesystems(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Thaw virtual machine file systems.
        This operation thaws a virtual machine's file systems using the QEMU guest agent when taking a live snapshot of a
        running virtual machine. Normally, this is done automatically by the manager, but this must be executed manually
        with the API for virtual machines using OpenStack Volume (Cinder) disks.
        Example:
        [source]
        ----
        POST /api/vms/123/thawfilesystems
        ----
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'thawfilesystems', None, headers, query, wait)

    def ticket(
        self,
        async=None,
        ticket=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Generates a time-sensitive authentication token for accessing a virtual machine's display.
        [source]
        ----
        POST /ovirt-engine/api/vms/123/ticket
        ----
        The client-provided action optionally includes a desired ticket value and/or an expiry time in seconds.
        In any case, the response specifies the actual ticket value and expiry used.
        [source,xml]
        ----
        <action>
          <ticket>
            <value>abcd12345</value>
            <expiry>120</expiry>
          </ticket>
        </action>
        ----
        [IMPORTANT]
        ====
        If the virtual machine is configured to support only one graphics protocol
        then the generated authentication token will be valid for that protocol.
        But if the virtual machine is configured to support multiple protocols,
        VNC and SPICE, then the authentication token will only be valid for
        the SPICE protocol.
        In order to obtain an authentication token for a specific protocol, for
        example for VNC, use the `ticket` method of the <<services/vm_graphics_console,
        service>> that manages the graphics consoles of the virtual machine, sending
        a request like this:
        [source]
        ----
        POST /ovirt-engine/api/vms/123/graphicsconsoles/456/ticket
        ----
        ====


        This method supports the following parameters:

        `async`:: Indicates if the generation of the ticket should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('ticket', ticket, types.Ticket),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            ticket=ticket,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'ticket', 'ticket', headers, query, wait)

    def undo_snapshot(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'undosnapshot', None, headers, query, wait)

    def update(
        self,
        vm,
        async=None,
        next_run=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('vm', vm, types.Vm),
            ('async', async, bool),
            ('next_run', next_run, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async
        if next_run is not None:
            next_run = Writer.render_boolean(next_run)
            query['next_run'] = next_run

        # Send the request and wait for the response:
        return self._internal_update(vm, headers, query, wait)

    def affinity_labels_service(self):
        """
        List of scheduling labels assigned to this VM.

        """
        return AssignedAffinityLabelsService(self._connection, '%s/affinitylabels' % self._path)

    def applications_service(self):
        """
        """
        return VmApplicationsService(self._connection, '%s/applications' % self._path)

    def cdroms_service(self):
        """
        """
        return VmCdromsService(self._connection, '%s/cdroms' % self._path)

    def disk_attachments_service(self):
        """
        List of disks attached to this virtual machine.

        """
        return DiskAttachmentsService(self._connection, '%s/diskattachments' % self._path)

    def graphics_consoles_service(self):
        """
        """
        return VmGraphicsConsolesService(self._connection, '%s/graphicsconsoles' % self._path)

    def host_devices_service(self):
        """
        """
        return VmHostDevicesService(self._connection, '%s/hostdevices' % self._path)

    def katello_errata_service(self):
        """
        Reference to the service that can show the applicable errata available on the virtual machine.
        This information is taken from Katello.

        """
        return KatelloErrataService(self._connection, '%s/katelloerrata' % self._path)

    def nics_service(self):
        """
        """
        return VmNicsService(self._connection, '%s/nics' % self._path)

    def numa_nodes_service(self):
        """
        """
        return VmNumaNodesService(self._connection, '%s/numanodes' % self._path)

    def permissions_service(self):
        """
        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def reported_devices_service(self):
        """
        """
        return VmReportedDevicesService(self._connection, '%s/reporteddevices' % self._path)

    def sessions_service(self):
        """
        Reference to the service that provides information about virtual machine user sessions.

        """
        return VmSessionsService(self._connection, '%s/sessions' % self._path)

    def snapshots_service(self):
        """
        """
        return SnapshotsService(self._connection, '%s/snapshots' % self._path)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def tags_service(self):
        """
        """
        return AssignedTagsService(self._connection, '%s/tags' % self._path)

    def watchdogs_service(self):
        """
        """
        return VmWatchdogsService(self._connection, '%s/watchdogs' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'affinitylabels':
            return self.affinity_labels_service()
        if path.startswith('affinitylabels/'):
            return self.affinity_labels_service().service(path[15:])
        if path == 'applications':
            return self.applications_service()
        if path.startswith('applications/'):
            return self.applications_service().service(path[13:])
        if path == 'cdroms':
            return self.cdroms_service()
        if path.startswith('cdroms/'):
            return self.cdroms_service().service(path[7:])
        if path == 'diskattachments':
            return self.disk_attachments_service()
        if path.startswith('diskattachments/'):
            return self.disk_attachments_service().service(path[16:])
        if path == 'graphicsconsoles':
            return self.graphics_consoles_service()
        if path.startswith('graphicsconsoles/'):
            return self.graphics_consoles_service().service(path[17:])
        if path == 'hostdevices':
            return self.host_devices_service()
        if path.startswith('hostdevices/'):
            return self.host_devices_service().service(path[12:])
        if path == 'katelloerrata':
            return self.katello_errata_service()
        if path.startswith('katelloerrata/'):
            return self.katello_errata_service().service(path[14:])
        if path == 'nics':
            return self.nics_service()
        if path.startswith('nics/'):
            return self.nics_service().service(path[5:])
        if path == 'numanodes':
            return self.numa_nodes_service()
        if path.startswith('numanodes/'):
            return self.numa_nodes_service().service(path[10:])
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'reporteddevices':
            return self.reported_devices_service()
        if path.startswith('reporteddevices/'):
            return self.reported_devices_service().service(path[16:])
        if path == 'sessions':
            return self.sessions_service()
        if path.startswith('sessions/'):
            return self.sessions_service().service(path[9:])
        if path == 'snapshots':
            return self.snapshots_service()
        if path.startswith('snapshots/'):
            return self.snapshots_service().service(path[10:])
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        if path == 'tags':
            return self.tags_service()
        if path.startswith('tags/'):
            return self.tags_service().service(path[5:])
        if path == 'watchdogs':
            return self.watchdogs_service()
        if path.startswith('watchdogs/'):
            return self.watchdogs_service().service(path[10:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmService:%s' % self._path


class VmApplicationService(Service):
    """
    A service that provides information about an application installed in a virtual machine.

    """

    def __init__(self, connection, path):
        super(VmApplicationService, self).__init__(connection, path)

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns the information about the application.


        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmApplicationService:%s' % self._path


class VmApplicationsService(Service):
    """
    A service that provides information about applications installed in a virtual machine.

    """

    def __init__(self, connection, path):
        super(VmApplicationsService, self).__init__(connection, path)
        self._application_service = None

    def list(
        self,
        filter=None,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns a list of applications installed in the virtual machine.


        This method supports the following parameters:

        `max`:: Sets the maximum number of applications to return. If not specified all the applications are returned.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def application_service(self, id):
        """
        Returns a reference to the service that provides information about a specific application.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmApplicationService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.application_service(path)
        return self.application_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmApplicationsService:%s' % self._path


class VmCdromService(Service):
    """
    Manages a CDROM device of a virtual machine.
    Changing and ejecting the disk is done using always the `update` method, to change the value of the `file`
    attribute.

    """

    def __init__(self, connection, path):
        super(VmCdromService, self).__init__(connection, path)

    def get(
        self,
        current=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns the information about this CDROM device.
        The information consists of `cdrom` attribute containing reference to the CDROM device, the virtual machine,
        and optionally the inserted disk.
        If there is a disk inserted then the `file` attribute will contain a reference to the ISO image:
        [source,xml]
        ----
        <cdrom href="..." id="00000000-0000-0000-0000-000000000000">
          <file id="mycd.iso"/>
          <vm href="/ovirt-engine/api/vms/123" id="123"/>
        </cdrom>
        ----
        If there is no disk inserted then the `file` attribute won't be reported:
        [source,xml]
        ----
        <cdrom href="..." id="00000000-0000-0000-0000-000000000000">
          <vm href="/ovirt-engine/api/vms/123" id="123"/>
        </cdrom>
        ----


        This method supports the following parameters:

        `current`:: Indicates if the operation should return the information for the currently running virtual machine. This
        parameter is optional, and the default value is `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('current', current, bool),
        ])

        # Build the URL:
        query = query or {}
        if current is not None:
            current = Writer.render_boolean(current)
            query['current'] = current

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def update(
        self,
        cdrom,
        current=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the information about this CDROM device.
        It allows to change or eject the disk by changing the value of the `file` attribute.
        For example, to insert or change the disk send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/vms/123/cdroms/00000000-0000-0000-0000-000000000000
        ----
        The body should contain the new value for the `file` attribute:
        [source,xml]
        ----
        <cdrom>
          <file id="mycd.iso"/>
        </cdrom>
        ----
        The value of the `id` attribute, `mycd.iso` in this example, should correspond to a file available in an
        attached ISO storage domain.
        To eject the disk use a `file` with an empty `id`:
        [source,xml]
        ----
        <cdrom>
          <file id=""/>
        </cdrom>
        ----
        By default the above operations change permanently the disk that will be visible to the virtual machine
        after the next boot, but they don't have any effect on the currently running virtual machine. If you want
        to change the disk that is visible to the current running virtual machine, add the `current=true` parameter.
        For example, to eject the current disk send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/vms/123/cdroms/00000000-0000-0000-0000-000000000000?current=true
        ----
        With a request body like this:
        [source,xml]
        ----
        <cdrom>
          <file id=""/>
        </cdrom>
        ----
        IMPORTANT: The changes made with the `current=true` parameter are never persisted, so they won't have any
        effect after the virtual machine is rebooted.


        This method supports the following parameters:

        `cdrom`:: The information about the CDROM device.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('cdrom', cdrom, types.Cdrom),
            ('current', current, bool),
        ])

        # Build the URL:
        query = query or {}
        if current is not None:
            current = Writer.render_boolean(current)
            query['current'] = current

        # Send the request and wait for the response:
        return self._internal_update(cdrom, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmCdromService:%s' % self._path


class VmCdromsService(Service):
    """
    Manages the CDROM devices of a virtual machine.
    Currently virtual machines have exactly one CDROM device. No new devices can be added, and the existing one can't
    be removed, thus there are no `add` or `remove` methods. Changing and ejecting CDROM disks is done with the
    <<services/vm_cdrom/methods/update, update>> method of the <<services/vm_cdrom, service>> that manages the
    CDROM device.

    """

    def __init__(self, connection, path):
        super(VmCdromsService, self).__init__(connection, path)
        self._cdrom_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns the list of CDROM devices of the virtual machine.


        This method supports the following parameters:

        `max`:: Sets the maximum number of CDROMs to return. If not specified all the CDROMs are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def cdrom_service(self, id):
        """
        Returns a reference to the service that manages a specific CDROM device.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmCdromService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.cdrom_service(path)
        return self.cdrom_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmCdromsService:%s' % self._path


class VmDiskService(MeasurableService):
    """
    """

    def __init__(self, connection, path):
        super(VmDiskService, self).__init__(connection, path)
        self._permissions_service = None
        self._statistics_service = None

    def activate(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the activation should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'activate', None, headers, query, wait)

    def deactivate(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the deactivation should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'deactivate', None, headers, query, wait)

    def export(
        self,
        async=None,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the export should be performed asynchronously.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('filter', filter, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            filter=filter,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'export', None, headers, query, wait)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def move(
        self,
        async=None,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the move should be performed asynchronously.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('filter', filter, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            filter=filter,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'move', None, headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Detach the disk from the virtual machine.
        NOTE: In version 3 of the API this used to also remove the disk completely from the system, but starting with
        version 4 it doesn't. If you need to remove it completely use the <<services/disk/methods/remove,remove
        method of the top level disk service>>.


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        disk,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('disk', disk, types.Disk),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(disk, headers, query, wait)

    def permissions_service(self):
        """
        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmDiskService:%s' % self._path


class VmDisksService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VmDisksService, self).__init__(connection, path)
        self._disk_service = None

    def add(
        self,
        disk,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('disk', disk, types.Disk),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(disk, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of disks to return. If not specified all the disks are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def disk_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmDiskService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.disk_service(path)
        return self.disk_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmDisksService:%s' % self._path


class VmGraphicsConsoleService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VmGraphicsConsoleService, self).__init__(connection, path)

    def get(
        self,
        current=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Gets graphics console configuration of the virtual machine.


        This method supports the following parameters:

        `current`:: Use the following query to obtain the current run-time configuration of the graphics console.
        [source]
        ----
        GET /ovit-engine/api/vms/123/graphicsconsoles/456?current=true
        ----
        The default value is `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('current', current, bool),
        ])

        # Build the URL:
        query = query or {}
        if current is not None:
            current = Writer.render_boolean(current)
            query['current'] = current

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def proxy_ticket(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the generation of the ticket should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'proxyticket', 'proxy_ticket', headers, query, wait)

    def remote_viewer_connection_file(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Generates the file which is compatible with `remote-viewer` client.
        Use the following request to generate remote viewer connection file of the graphics console.
        Note that this action generates the file only if virtual machine is running.
        [source]
        ----
        POST /ovirt-engine/api/vms/123/graphicsconsoles/456/remoteviewerconnectionfile
        ----
        The `remoteviewerconnectionfile` action does not take any action specific parameters,
        so the request body should contain an empty `action`:
        [source,xml]
        ----
        <action/>
        ----
        The response contains the file, which can be used with `remote-viewer` client.
        [source,xml]
        ----
        <action>
          <remote_viewer_connection_file>
            [virt-viewer]
            type=spice
            host=192.168.1.101
            port=-1
            password=123456789
            delete-this-file=1
            fullscreen=0
            toggle-fullscreen=shift+f11
            release-cursor=shift+f12
            secure-attention=ctrl+alt+end
            tls-port=5900
            enable-smartcard=0
            enable-usb-autoshare=0
            usb-filter=null
            tls-ciphers=DEFAULT
            host-subject=O=local,CN=example.com
            ca=...
          </remote_viewer_connection_file>
        </action>
        ----
        E.g., to fetch the content of remote viewer connection file and save it into temporary file, user can use
        oVirt Python SDK as follows:
        [source,python]
        ----
        # Find the virtual machine:
        vm = vms_service.list(search='name=myvm')[0]
        # Locate the service that manages the virtual machine, as that is where
        # the locators are defined:
        vm_service = vms_service.vm_service(vm.id)
        # Find the graphic console of the virtual machine:
        graphics_consoles_service = vm_service.graphics_consoles_service()
        graphics_console = graphics_consoles_service.list()[0]
        # Generate the remote viewer connection file:
        console_service = graphics_consoles_service.console_service(graphics_console.id)
        remote_viewer_connection_file = console_service.remote_viewer_connection_file()
        # Write the content to file "/tmp/remote_viewer_connection_file.vv"
        path = "/tmp/remote_viewer_connection_file.vv"
        with open(path, "w") as f:
            f.write(remote_viewer_connection_file)
        ----
        When you create the remote viewer connection file, then you can connect to virtual machine graphic console,
        as follows:
        [source,bash]
        ----
        #!/bin/sh -ex
        remote-viewer --ovirt-ca-file=/etc/pki/ovirt-engine/ca.pem /tmp/remote_viewer_connection_file.vv
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'remoteviewerconnectionfile', 'remote_viewer_connection_file', headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove the graphics console from the virtual machine.


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def ticket(
        self,
        ticket=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Generates a time-sensitive authentication token for accessing this virtual machine's console.
        [source]
        ----
        POST /ovirt-engine/api/vms/123/graphicsconsoles/456/ticket
        ----
        The client-provided action optionally includes a desired ticket value and/or an expiry time in seconds.
        In any case, the response specifies the actual ticket value and expiry used.
        [source,xml]
        ----
        <action>
          <ticket>
            <value>abcd12345</value>
            <expiry>120</expiry>
          </ticket>
        </action>
        ----


        This method supports the following parameters:

        `ticket`:: The generated ticket that can be used to access this console.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('ticket', ticket, types.Ticket),
        ])

        # Populate the action:
        action = types.Action(
            ticket=ticket,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'ticket', 'ticket', headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmGraphicsConsoleService:%s' % self._path


class VmGraphicsConsolesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VmGraphicsConsolesService, self).__init__(connection, path)
        self._console_service = None

    def add(
        self,
        console,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add new graphics console to the virtual machine.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('console', console, types.GraphicsConsole),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(console, headers, query, wait)

    def list(
        self,
        current=None,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists all the configured graphics consoles of the virtual machine.


        This method supports the following parameters:

        `max`:: Sets the maximum number of consoles to return. If not specified all the consoles are returned.

        `current`:: Use the following query to obtain the current run-time configuration of the graphics consoles.
        [source]
        ----
        GET /ovirt-engine/api/vms/123/graphicsconsoles?current=true
        ----
        The default value is `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('current', current, bool),
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if current is not None:
            current = Writer.render_boolean(current)
            query['current'] = current
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def console_service(self, id):
        """
        Returns a reference to the service that manages a specific virtual machine graphics console.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmGraphicsConsoleService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.console_service(path)
        return self.console_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmGraphicsConsolesService:%s' % self._path


class VmHostDeviceService(Service):
    """
    A service to manage individual host device attached to a virtual machine.

    """

    def __init__(self, connection, path):
        super(VmHostDeviceService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieve information about particular host device attached to given virtual machine.
        Example:
        [source]
        ----
        GET /ovirt-engine/api/vms/123/hostdevices/456
        ----
        [source,xml]
        ----
        <host_device href="/ovirt-engine/api/hosts/543/devices/456" id="456">
          <name>pci_0000_04_00_0</name>
          <capability>pci</capability>
          <iommu_group>30</iommu_group>
          <placeholder>true</placeholder>
          <product id="0x13ba">
            <name>GM107GL [Quadro K2200]</name>
          </product>
          <vendor id="0x10de">
            <name>NVIDIA Corporation</name>
          </vendor>
          <host href="/ovirt-engine/api/hosts/543" id="543"/>
          <parent_device href="/ovirt-engine/api/hosts/543/devices/456" id="456">
            <name>pci_0000_00_03_0</name>
          </parent_device>
          <vm href="/ovirt-engine/api/vms/123" id="123"/>
        </host_device>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove the attachment of this host device from given virtual machine.
        NOTE: In case this device serves as an IOMMU placeholder, it cannot be removed (remove will result only
        in setting its `placeholder` flag to `true`). Note that all IOMMU placeholder devices will be removed
        automatically as soon as there will be no more non-placeholder devices (all devices from given IOMMU
        group are detached).
        [source]
        ----
        DELETE /ovirt-engine/api/vms/123/hostdevices/456
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmHostDeviceService:%s' % self._path


class VmHostDevicesService(Service):
    """
    A service to manage host devices attached to a virtual machine.

    """

    def __init__(self, connection, path):
        super(VmHostDevicesService, self).__init__(connection, path)
        self._device_service = None

    def add(
        self,
        device,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Attach target device to given virtual machine.
        Example:
        [source]
        ----
        POST /ovirt-engine/api/vms/123/hostdevices
        ----
        With request body of type <<types/host_device,HostDevice>>, for example
        [source,xml]
        ----
        <host_device id="123" />
        ----
        NOTE: A necessary precondition for a successful host device attachment is that the virtual machine must be pinned
        to *exactly* one host. The device ID is then taken relative to this host.
        NOTE: Attachment of a PCI device that is part of a bigger IOMMU group will result in attachment of the remaining
        devices from that IOMMU group as "placeholders". These devices are then identified using the `placeholder`
        attribute of the <<types/host_device,HostDevice>> type set to `true`.
        In case you want attach a device that already serves as an IOMMU placeholder, simply issue an explicit Add operation
        for it, and its `placeholder` flag will be cleared, and the device will be accessible to the virtual machine.


        This method supports the following parameters:

        `device`:: The host device to be attached to given virtual machine.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('device', device, types.HostDevice),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(device, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List the host devices assigned to given virtual machine.


        This method supports the following parameters:

        `max`:: Sets the maximum number of devices to return. If not specified all the devices are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def device_service(self, id):
        """
        Returns a reference to the service that manages a specific host device attached to given virtual machine.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmHostDeviceService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.device_service(path)
        return self.device_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmHostDevicesService:%s' % self._path


class VmNicService(MeasurableService):
    """
    """

    def __init__(self, connection, path):
        super(VmNicService, self).__init__(connection, path)
        self._network_filter_parameters_service = None
        self._reported_devices_service = None
        self._statistics_service = None

    def activate(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the activation should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'activate', None, headers, query, wait)

    def deactivate(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the deactivation should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'deactivate', None, headers, query, wait)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the NIC.
        For example, to remove the NIC with id `456` from the virtual machine with id `123` send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/vms/123/nics/456
        ----
        [IMPORTANT]
        ====
        The hotplugging feature only supports virtual machine operating systems with hotplugging operations.
        Example operating systems include:
        - Red Hat Enterprise Linux 6
        - Red Hat Enterprise Linux 5
        - Windows Server 2008 and
        - Windows Server 2003
        ====


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        nic,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the NIC.
        For example, to update the NIC having with `456` belonging to virtual the machine with id `123` send a request
        like this:
        [source]
        ----
        PUT /ovirt-engine/api/vms/123/nics/456
        ----
        With a request body like this:
        [source,xml]
        ----
        <nic>
          <name>mynic</name>
          <interface>e1000</interface>
          <vnic_profile id='789'/>
        </nic>
        ----
        [IMPORTANT]
        ====
        The hotplugging feature only supports virtual machine operating systems with hotplugging operations.
        Example operating systems include:
        - Red Hat Enterprise Linux 6
        - Red Hat Enterprise Linux 5
        - Windows Server 2008 and
        - Windows Server 2003
        ====


        """
        # Check the types of the parameters:
        Service._check_types([
            ('nic', nic, types.Nic),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(nic, headers, query, wait)

    def network_filter_parameters_service(self):
        """
        Reference to the service that manages the network filter parameters of the NIC.
        A single top-level network filter may assigned to the NIC by the NIC's <<types/vnic_profile,vNIC Profile>>.

        """
        return NetworkFilterParametersService(self._connection, '%s/networkfilterparameters' % self._path)

    def reported_devices_service(self):
        """
        """
        return VmReportedDevicesService(self._connection, '%s/reporteddevices' % self._path)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'networkfilterparameters':
            return self.network_filter_parameters_service()
        if path.startswith('networkfilterparameters/'):
            return self.network_filter_parameters_service().service(path[24:])
        if path == 'reporteddevices':
            return self.reported_devices_service()
        if path.startswith('reporteddevices/'):
            return self.reported_devices_service().service(path[16:])
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmNicService:%s' % self._path


class VmNicsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VmNicsService, self).__init__(connection, path)
        self._nic_service = None

    def add(
        self,
        nic,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adds a NIC to the virtual machine.
        The following example adds a network interface named `mynic` using `virtio` and the `ovirtmgmt` network to the
        virtual machine.
        [source]
        ----
        POST /ovirt-engine/api/vms/123/nics
        ----
        [source,xml]
        ----
        <nic>
          <interface>virtio</interface>
          <name>mynic</name>
          <network>
            <name>ovirtmgmt</name>
          </network>
        </nic>
        ----
        The following example sends that request using `curl`:
        [source,bash]
        ----
        curl \
        --request POST \
        --header "Version: 4" \
        --header "Content-Type: application/xml" \
        --header "Accept: application/xml" \
        --user "admin@internal:mypassword" \
        --cacert /etc/pki/ovirt-engine/ca.pem \
        --data '
        <nic>
          <name>mynic</name>
          <network>
            <name>ovirtmgmt</name>
          </network>
        </nic>
        ' \
        https://myengine.example.com/ovirt-engine/api/vms/123/nics
        ----
        [IMPORTANT]
        ====
        The hotplugging feature only supports virtual machine operating systems with hotplugging operations.
        Example operating systems include:
        - Red Hat Enterprise Linux 6
        - Red Hat Enterprise Linux 5
        - Windows Server 2008 and
        - Windows Server 2003
        ====


        """
        # Check the types of the parameters:
        Service._check_types([
            ('nic', nic, types.Nic),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(nic, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of NICs to return. If not specified all the NICs are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def nic_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmNicService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.nic_service(path)
        return self.nic_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmNicsService:%s' % self._path


class VmNumaNodeService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VmNumaNodeService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a virtual NUMA node.
        An example of removing a virtual NUMA node:
        [source]
        ----
        DELETE /ovirt-engine/api/vms/123/numanodes/456
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        node,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates a virtual NUMA node.
        An example of pinning a virtual NUMA node to a physical NUMA node on the host:
        [source]
        ----
        PUT /ovirt-engine/api/vms/123/numanodes/456
        ----
        The request body should contain the following:
        [source,xml]
        ----
        <vm_numa_node>
          <numa_node_pins>
            <numa_node_pin>
              <index>0</index>
            </numa_node_pin>
          </numa_node_pins>
        </vm_numa_node>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('node', node, types.VirtualNumaNode),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(node, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmNumaNodeService:%s' % self._path


class VmNumaNodesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VmNumaNodesService, self).__init__(connection, path)
        self._node_service = None

    def add(
        self,
        node,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new virtual NUMA node for the virtual machine.
        An example of creating a NUMA node:
        [source]
        ----
        POST /ovirt-engine/api/vms/c7ecd2dc/numanodes
        Accept: application/xml
        Content-type: application/xml
        ----
        The request body can contain the following:
        [source,xml]
        ----
        <vm_numa_node>
          <cpu>
            <cores>
              <core>
                <index>0</index>
              </core>
            </cores>
          </cpu>
          <index>0</index>
          <memory>1024</memory>
        </vm_numa_node>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('node', node, types.VirtualNumaNode),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(node, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists virtual NUMA nodes of a virtual machine.


        This method supports the following parameters:

        `max`:: Sets the maximum number of nodes to return. If not specified all the nodes are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def node_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmNumaNodeService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.node_service(path)
        return self.node_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmNumaNodesService:%s' % self._path


class VmPoolService(Service):
    """
    A service to manage a virtual machines pool.

    """

    def __init__(self, connection, path):
        super(VmPoolService, self).__init__(connection, path)
        self._permissions_service = None

    def allocate_vm(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation allocates a virtual machine in the virtual machine pool.
        [source]
        ----
        POST /ovirt-engine/api/vmpools/123/allocatevm
        ----
        The allocate virtual machine action does not take any action specific parameters, so the request body should
        contain an empty `action`:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the allocation should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'allocatevm', None, headers, query, wait)

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get the virtual machine pool.
        [source]
        ----
        GET /ovirt-engine/api/vmpools/123
        ----
        You will get a XML response like that one:
        [source,xml]
        ----
        <vm_pool id="123">
          <actions>...</actions>
          <name>MyVmPool</name>
          <description>MyVmPool description</description>
          <link href="/ovirt-engine/api/vmpools/123/permissions" rel="permissions"/>
          <max_user_vms>1</max_user_vms>
          <prestarted_vms>0</prestarted_vms>
          <size>100</size>
          <stateful>false</stateful>
          <type>automatic</type>
          <use_latest_template_version>false</use_latest_template_version>
          <cluster id="123"/>
          <template id="123"/>
          <vm id="123">...</vm>
          ...
        </vm_pool>
        ----


        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a virtual machine pool.
        [source]
        ----
        DELETE /ovirt-engine/api/vmpools/123
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        pool,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Update the virtual machine pool.
        [source]
        ----
        PUT /ovirt-engine/api/vmpools/123
        ----
        The `name`, `description`, `size`, `prestarted_vms` and `max_user_vms`
        attributes can be updated after the virtual machine pool has been
        created.
        [source,xml]
        ----
        <vmpool>
          <name>VM_Pool_B</name>
          <description>Virtual Machine Pool B</description>
          <size>3</size>
          <prestarted_vms>1</size>
          <max_user_vms>2</size>
        </vmpool>
        ----


        This method supports the following parameters:

        `pool`:: The virtual machine pool that is being updated.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('pool', pool, types.VmPool),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(pool, headers, query, wait)

    def permissions_service(self):
        """
        Reference to a service managing the virtual machine pool assigned permissions.

        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmPoolService:%s' % self._path


class VmPoolsService(Service):
    """
    Provides read-write access to virtual machines pools.

    """

    def __init__(self, connection, path):
        super(VmPoolsService, self).__init__(connection, path)
        self._pool_service = None

    def add(
        self,
        pool,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new virtual machine pool.
        A new pool requires the `name`, `cluster` and `template` attributes. Identify the cluster and template with the
        `id` or `name` nested attributes:
        [source]
        ----
        POST /ovirt-engine/api/vmpools
        ----
        With the following body:
        [source,xml]
        ----
        <vmpool>
          <name>mypool</name>
          <cluster id="123"/>
          <template id="456"/>
        </vmpool>
        ----


        This method supports the following parameters:

        `pool`:: Pool to add.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('pool', pool, types.VmPool),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(pool, headers, query, wait)

    def list(
        self,
        case_sensitive=None,
        filter=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get a list of available virtual machines pools.
        [source]
        ----
        GET /ovirt-engine/api/vmpools
        ----
        You will receive the following response:
        [source,xml]
        ----
        <vm_pools>
          <vm_pool id="123">
            ...
          </vm_pool>
          ...
        </vm_pools>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of pools to return. If this value is not specified, all of the pools are returned.

        `search`:: A query string used to restrict the returned pools.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('case_sensitive', case_sensitive, bool),
            ('filter', filter, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def pool_service(self, id):
        """
        Reference to the service that manages a specific virtual machine pool.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmPoolService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.pool_service(path)
        return self.pool_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmPoolsService:%s' % self._path


class VmReportedDeviceService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VmReportedDeviceService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmReportedDeviceService:%s' % self._path


class VmReportedDevicesService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VmReportedDevicesService, self).__init__(connection, path)
        self._reported_device_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of devices to return. If not specified all the devices are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def reported_device_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmReportedDeviceService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.reported_device_service(path)
        return self.reported_device_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmReportedDevicesService:%s' % self._path


class VmSessionService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VmSessionService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmSessionService:%s' % self._path


class VmSessionsService(Service):
    """
    Provides information about virtual machine user sessions.

    """

    def __init__(self, connection, path):
        super(VmSessionsService, self).__init__(connection, path)
        self._session_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Lists all user sessions for this virtual machine.
        For example, to retrieve the session information for virtual machine `123` send a request like this:
        [source]
        ----
        GET /ovirt-engine/api/vms/123/sessions
        ----
        The response body will contain something like this:
        [source,xml]
        ----
        <sessions>
          <session href="/ovirt-engine/api/vms/123/sessions/456" id="456">
            <console_user>true</console_user>
            <ip>
              <address>192.168.122.1</address>
            </ip>
            <user href="/ovirt-engine/api/users/789" id="789"/>
            <vm href="/ovirt-engine/api/vms/123" id="123"/>
          </session>
          ...
        </sessions>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of sessions to return. If not specified all the sessions are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def session_service(self, id):
        """
        Reference to the service that manages a specific session.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmSessionService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.session_service(path)
        return self.session_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmSessionsService:%s' % self._path


class VmWatchdogService(Service):
    """
    A service managing a watchdog on virtual machines.

    """

    def __init__(self, connection, path):
        super(VmWatchdogService, self).__init__(connection, path)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Returns the information about the watchdog.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the watchdog from the virtual machine.
        For example, to remove a watchdog from a virtual machine, send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/vms/123/watchdogs/00000000-0000-0000-0000-000000000000
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        watchdog,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the information about the watchdog.
        You can update the information using `action` and `model` elements.
        For example, to update a watchdog, send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/vms/123/watchdogs
        <watchdog>
          <action>reset</action>
        </watchdog>
        ----
        with response body:
        [source,xml]
        ----
        <watchdog href="/ovirt-engine/api/vms/123/watchdogs/00000000-0000-0000-0000-000000000000" id="00000000-0000-0000-0000-000000000000">
          <vm href="/ovirt-engine/api/vms/123" id="123"/>
          <action>reset</action>
          <model>i6300esb</model>
        </watchdog>
        ----


        This method supports the following parameters:

        `watchdog`:: The information about the watchdog.
        The request data must contain at least one of `model` and `action`
        elements. The response data contains complete information about the
        updated watchdog.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('watchdog', watchdog, types.Watchdog),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(watchdog, headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VmWatchdogService:%s' % self._path


class VmWatchdogsService(Service):
    """
    Lists the watchdogs of a virtual machine.

    """

    def __init__(self, connection, path):
        super(VmWatchdogsService, self).__init__(connection, path)
        self._watchdog_service = None

    def add(
        self,
        watchdog,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Adds new watchdog to the virtual machine.
        For example, to add a watchdog to a virtual machine, send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/vms/123/watchdogs
        <watchdog>
          <action>poweroff</action>
          <model>i6300esb</model>
        </watchdog>
        ----
        with response body:
        [source,xml]
        ----
        <watchdog href="/ovirt-engine/api/vms/123/watchdogs/00000000-0000-0000-0000-000000000000" id="00000000-0000-0000-0000-000000000000">
          <vm href="/ovirt-engine/api/vms/123" id="123"/>
          <action>poweroff</action>
          <model>i6300esb</model>
        </watchdog>
        ----


        This method supports the following parameters:

        `watchdog`:: The information about the watchdog.
        The request data must contain `model` element (such as `i6300esb`) and `action` element
        (one of `none`, `reset`, `poweroff`, `dump`, `pause`). The response data additionally
        contains references to the added watchdog and to the virtual machine.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('watchdog', watchdog, types.Watchdog),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(watchdog, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        The list of watchdogs of the virtual machine.


        This method supports the following parameters:

        `max`:: Sets the maximum number of watchdogs to return. If not specified all the watchdogs are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def watchdog_service(self, id):
        """
        Returns a reference to the service that manages a specific watchdog.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmWatchdogService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.watchdog_service(path)
        return self.watchdog_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmWatchdogsService:%s' % self._path


class VmsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(VmsService, self).__init__(connection, path)
        self._vm_service = None

    def add(
        self,
        vm,
        clone=None,
        clone_permissions=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Creates a new virtual machine.
        The virtual machine can be created in different ways:
        - From a template. In this case the identifier or name of the template must be provided. For example, using a
          plain shell script and XML:
        [source,bash]
        ----
        #!/bin/sh -ex
        url="https://engine.example.com/ovirt-engine/api"
        user="admin@internal"
        password="..."
        curl \
        --verbose \
        --cacert /etc/pki/ovirt-engine/ca.pem \
        --user "${user}:${password}" \
        --request POST \
        --header "Version: 4" \
        --header "Content-Type: application/xml" \
        --header "Accept: application/xml" \
        --data '
        <vm>
          <name>myvm</name>
          <template>
            <name>Blank</name>
          </template>
          <cluster>
            <name>mycluster</name>
          </cluster>
        </vm>
        ' \
        "${url}/vms"
        ----
        - From a snapshot. In this case the identifier of the snapshot has to be provided. For example, using a plain
          shel script and XML:
        [source,bash]
        ----
        #!/bin/sh -ex
        url="https://engine.example.com/ovirt-engine/api"
        user="admin@internal"
        password="..."
        curl \
        --verbose \
        --cacert /etc/pki/ovirt-engine/ca.pem \
        --user "${user}:${password}" \
        --request POST \
        --header "Content-Type: application/xml" \
        --header "Accept: application/xml" \
        --data '
        <vm>
          <name>myvm</name>
          <snapshots>
            <snapshot id="266742a5-6a65-483c-816d-d2ce49746680"/>
          </snapshots>
          <cluster>
            <name>mycluster</name>
          </cluster>
        </vm>
        ' \
        "${url}/vms"
        ----
        When creating a virtual machine from a template or from a snapshot it is usually useful to explicitly indicate
        in what storage domain to create the disks for the virtual machine. If the virtual machine is created from
        a template then this is achieved passing a set of `disk_attachment` elements that indicate the mapping:
        [source,xml]
        ----
        <vm>
          ...
          <disk_attachments>
            <disk_attachment>
              <disk id="8d4bd566-6c86-4592-a4a7-912dbf93c298">
                <storage_domains>
                  <storage_domain id="9cb6cb0a-cf1d-41c2-92ca-5a6d665649c9"/>
                </storage_domains>
              </disk>
            <disk_attachment>
          </disk_attachments>
        </vm>
        ----
        When the virtual machine is created from a snapshot this set of disks is slightly different, it uses the
        `image_id` attribute instead of `id`.
        [source,xml]
        ----
        <vm>
          ...
          <disk_attachments>
            <disk_attachment>
              <disk>
                <image_id>8d4bd566-6c86-4592-a4a7-912dbf93c298</image_id>
                <storage_domains>
                  <storage_domain id="9cb6cb0a-cf1d-41c2-92ca-5a6d665649c9"/>
                </storage_domains>
              </disk>
            <disk_attachment>
          </disk_attachments>
        </vm>
        ----
        It is possible to specify additional virtual machine parameters in the XML description, e.g. a virtual machine
        of `desktop` type, with 2 GiB of RAM and additional description can be added sending a request body like the
        following:
        [source,xml]
        ----
        <vm>
          <name>myvm</name>
          <description>My Desktop Virtual Machine</description>
          <type>desktop</type>
          <memory>2147483648</memory>
          ...
        </vm>
        ----
        A bootable CDROM device can be set like this:
        [source,xml]
        ----
        <vm>
          ...
          <os>
            <boot dev="cdrom"/>
          </os>
        </vm>
        ----
        In order to boot from CDROM, you first need to insert a disk, as described in the
        <<services/vm_cdrom, CDROM service>>. Then booting from that CDROM can be specified using the `os.boot.devices`
        attribute:
        [source,xml]
        ----
        <vm>
          ...
          <os>
            <boot>
              <devices>
                <device>cdrom</device>
              </devices>
            </boot>
          </os>
        </vm>
        ----
        In all cases the name or identifier of the cluster where the virtual machine will be created is mandatory.


        """
        # Check the types of the parameters:
        Service._check_types([
            ('vm', vm, types.Vm),
            ('clone', clone, bool),
            ('clone_permissions', clone_permissions, bool),
        ])

        # Build the URL:
        query = query or {}
        if clone is not None:
            clone = Writer.render_boolean(clone)
            query['clone'] = clone
        if clone_permissions is not None:
            clone_permissions = Writer.render_boolean(clone_permissions)
            query['clone_permissions'] = clone_permissions

        # Send the request and wait for the response:
        return self._internal_add(vm, headers, query, wait)

    def list(
        self,
        all_content=None,
        case_sensitive=None,
        filter=None,
        max=None,
        search=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `search`:: A query string used to restrict the returned virtual machines.

        `max`:: The maximum number of results to return.

        `case_sensitive`:: Indicates if the search performed using the `search` parameter should be performed taking case into
        account. The default value is `true`, which means that case is taken into account. If you want to search
        ignoring case set it to `false`.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `all_content`:: Indicates if all the attributes of the virtual machines should be included in the response.
        By default the following attributes are excluded:
        - `console`
        - `initialization.configuration.data` - The OVF document describing the virtual machine.
        - `rng_source`
        - `soundcard`
        - `virtio_scsi`
        For example, to retrieve the complete representation of the virtual machines send a request like this:
        ....
        GET /ovirt-engine/api/vms?all_content=true
        ....
        NOTE: The reason for not including these attributes is performance: they are seldom used and they require
        additional queries to the database. So try to use the this parameter only when it is really needed.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('all_content', all_content, bool),
            ('case_sensitive', case_sensitive, bool),
            ('filter', filter, bool),
            ('max', max, int),
            ('search', search, str),
        ])

        # Build the URL:
        query = query or {}
        if all_content is not None:
            all_content = Writer.render_boolean(all_content)
            query['all_content'] = all_content
        if case_sensitive is not None:
            case_sensitive = Writer.render_boolean(case_sensitive)
            query['case_sensitive'] = case_sensitive
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max
        if search is not None:
            query['search'] = search

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def vm_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return VmService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.vm_service(path)
        return self.vm_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VmsService:%s' % self._path


class VnicProfileService(Service):
    """
    This service manages a vNIC profile.

    """

    def __init__(self, connection, path):
        super(VnicProfileService, self).__init__(connection, path)
        self._permissions_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves details about a vNIC profile.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the vNIC profile.


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def update(
        self,
        profile,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates details of a vNIC profile.


        This method supports the following parameters:

        `profile`:: The vNIC profile that is being updated.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('profile', profile, types.VnicProfile),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(profile, headers, query, wait)

    def permissions_service(self):
        """
        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'VnicProfileService:%s' % self._path


class VnicProfilesService(Service):
    """
    This service manages the collection of all vNIC profiles.

    """

    def __init__(self, connection, path):
        super(VnicProfilesService, self).__init__(connection, path)
        self._profile_service = None

    def add(
        self,
        profile,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Add a vNIC profile.
        For example to add vNIC profile `123` to network `456` send a request to:
        [source]
        ----
        POST /ovirt-engine/api/networks/456/vnicprofiles
        ----
        With the following body:
        [source,xml]
        ----
        <vnic_profile id="123">
          <name>new_vNIC_name</name>
          <pass_through>
            <mode>disabled</mode>
          </pass_through>
          <port_mirroring>false</port_mirroring>
        </vnic_profile>
        ----
        Please note that there is a default network filter to each VNIC profile.
        For more details of how the default network filter is calculated please refer to
        the documentation in <<services/network_filters,NetworkFilters>>.
        The output of creating a new VNIC profile depends in the  body  arguments that were given.
        In case no network filter was given, the default network filter will be configured. For example:
        [source,xml]
        ----
        <vnic_profile href="/ovirt-engine/api/vnicprofiles/123" id="123">
          <name>new_vNIC_name</name>
          <link href="/ovirt-engine/api/vnicprofiles/123/permissions" rel="permissions"/>
          <pass_through>
            <mode>disabled</mode>
          </pass_through>
          <port_mirroring>false</port_mirroring>
          <network href="/ovirt-engine/api/networks/456" id="456"/>
          <network_filter href="/ovirt-engine/api/networkfilters/789" id="789"/>
        </vnic_profile>
        ----
        In case an empty network filter was given, no network filter will be configured for the specific VNIC profile
        regardless of the VNIC profile's default network filter. For example:
        [source,xml]
        ----
        <vnic_profile>
          <name>no_network_filter</name>
          <network_filter/>
        </vnic_profile>
        ----
        In case that a specific valid network filter id was given, the VNIC profile will be configured with the given
        network filter regardless of the VNIC profiles's default network filter. For example:
        [source,xml]
        ----
        <vnic_profile>
          <name>user_choice_network_filter</name>
          <network_filter id= "0000001b-001b-001b-001b-0000000001d5"/>
        </vnic_profile>
        ----


        This method supports the following parameters:

        `profile`:: The vNIC profile that is being added.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('profile', profile, types.VnicProfile),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(profile, headers, query, wait)

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        List all vNIC profiles.


        This method supports the following parameters:

        `max`:: Sets the maximum number of profiles to return. If not specified all the profiles are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def profile_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return VnicProfileService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.profile_service(path)
        return self.profile_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'VnicProfilesService:%s' % self._path


class WeightService(Service):
    """
    """

    def __init__(self, connection, path):
        super(WeightService, self).__init__(connection, path)

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'WeightService:%s' % self._path


class WeightsService(Service):
    """
    """

    def __init__(self, connection, path):
        super(WeightsService, self).__init__(connection, path)
        self._weight_service = None

    def add(
        self,
        weight,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('weight', weight, types.Weight),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_add(weight, headers, query, wait)

    def list(
        self,
        filter=None,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `max`:: Sets the maximum number of weights to return. If not specified all the weights are returned.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def weight_service(self, id):
        """
        """
        Service._check_types([
            ('id', id, str),
        ])
        return WeightService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.weight_service(path)
        return self.weight_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'WeightsService:%s' % self._path


class AttachedStorageDomainDiskService(MeasurableService):
    """
    Manages a single disk available in a storage domain attached to a data center.
    IMPORTANT: Since version 4.2 of the engine this service is intended only to list disks available in the storage
    domain, and to register unregistered disks. All the other operations, like copying a disk, moving a disk, etc, have
    been deprecated and will be removed in the future. To perform those operations use the <<services/disks, service
    that manages all the disks of the system>>, or the <<services/disk, service that manages an specific disk>>.

    """

    def __init__(self, connection, path):
        super(AttachedStorageDomainDiskService, self).__init__(connection, path)
        self._permissions_service = None
        self._statistics_service = None

    def copy(
        self,
        disk=None,
        storage_domain=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Copies a disk to the specified storage domain.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To copy a disk use the <<services/disk/methods/copy, copy>>
        operation of the service that manages that disk.


        This method supports the following parameters:

        `disk`:: Description of the resulting disk.

        `storage_domain`:: The storage domain where the new disk will be created.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('disk', disk, types.Disk),
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Populate the action:
        action = types.Action(
            disk=disk,
            storage_domain=storage_domain,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'copy', None, headers, query, wait)

    def export(
        self,
        storage_domain=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Exports a disk to an export storage domain.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To export a disk use the <<services/disk/methods/export, export>>
        operation of the service that manages that disk.


        This method supports the following parameters:

        `storage_domain`:: The export storage domain where the disk should be exported to.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Populate the action:
        action = types.Action(
            storage_domain=storage_domain,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'export', None, headers, query, wait)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the description of the disk.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def move(
        self,
        async=None,
        filter=None,
        storage_domain=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Moves a disk to another storage domain.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To move a disk use the <<services/disk/methods/move, move>>
        operation of the service that manages that disk.


        This method supports the following parameters:

        `storage_domain`:: The storage domain where the disk will be moved to.

        `async`:: Indicates if the move should be performed asynchronously.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('filter', filter, bool),
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            filter=filter,
            storage_domain=storage_domain,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'move', None, headers, query, wait)

    def register(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Registers an unregistered disk.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'register', None, headers, query, wait)

    def remove(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a disk.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To remove a disk use the <<services/disk/methods/remove, remove>>
        operation of the service that manages that disk.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def sparsify(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Sparsify the disk.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To remove a disk use the <<services/disk/methods/remove, remove>>
        operation of the service that manages that disk.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'sparsify', None, headers, query, wait)

    def update(
        self,
        disk,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Updates the disk.
        IMPORTANT: Since version 4.2 of the engine this operation is deprecated, and preserved only for backwards
        compatibility. It will be removed in the future. To update a disk use the
        <<services/disk/methods/update, update>> operation of the service that manages that disk.


        This method supports the following parameters:

        `disk`:: The update to apply to the disk.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('disk', disk, types.Disk),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_update(disk, headers, query, wait)

    def permissions_service(self):
        """
        Reference to the service that manages the permissions assigned to the disk.

        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'AttachedStorageDomainDiskService:%s' % self._path


class DiskService(MeasurableService):
    """
    Manages a single disk.

    """

    def __init__(self, connection, path):
        super(DiskService, self).__init__(connection, path)
        self._permissions_service = None
        self._statistics_service = None

    def copy(
        self,
        async=None,
        disk=None,
        filter=None,
        storage_domain=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation copies a disk to the specified storage domain.
        For example, copy of a disk can be facilitated using the following request:
        [source]
        ----
        POST /ovirt-engine/api/disks/123/copy
        ----
        With a request body like this:
        [source,xml]
        ----
        <action>
          <storage_domain id="456"/>
          <disk>
            <name>mydisk</name>
          </disk>
        </action>
        ----


        This method supports the following parameters:

        `disk`:: Description of the resulting disk. The only accepted value is the `name` attribute, which will be the name
        used for the new disk. For example, to copy disk `123` using `myname` as the name for the new disk, send
        a request like this:
        ....
        POST /ovirt-engine/disks/123
        ....
        With a request body like this:
        [source,xml]
        ----
        <action>
          <disk>
            <name>mydisk<name>
          </disk>
          <storage_domain id="456"/>
        </action>
        ----

        `storage_domain`:: The storage domain where the new disk will be created. Can be specified using the `id` or `name`
        attributes. For example, to copy a disk to the storage domain named `mydata` send a request like this:
        ....
        POST /ovirt-engine/api/storagedomains/123/disks/789
        ....
        With a request body like this:
        [source,xml]
        ----
        <action>
          <storage_domain>
            <name>mydata</name>
          </storage_domain>
        </action>
        ----

        `async`:: Indicates if the copy should be performed asynchronously.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('disk', disk, types.Disk),
            ('filter', filter, bool),
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            disk=disk,
            filter=filter,
            storage_domain=storage_domain,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'copy', None, headers, query, wait)

    def export(
        self,
        async=None,
        filter=None,
        storage_domain=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Exports a disk to an export storage domain.


        This method supports the following parameters:

        `storage_domain`:: The export storage domain where the disk should be exported to.

        `async`:: Indicates if the export should be performed asynchronously.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('filter', filter, bool),
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            filter=filter,
            storage_domain=storage_domain,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'export', None, headers, query, wait)

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the description of the disk.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def move(
        self,
        async=None,
        filter=None,
        storage_domain=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Moves a disk to another storage domain.
        For example, to move the disk with identifier `123` to a storage domain with identifier `456` send the following
        request:
        [source]
        ----
        POST /ovirt-engine/api/disks/123/move
        ----
        With the following request body:
        [source,xml]
        ----
        <action>
          <storage_domain id="456"/>
        </action>
        ----


        This method supports the following parameters:

        `storage_domain`:: The storage domain where the disk will be moved to.

        `async`:: Indicates if the move should be performed asynchronously.

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('filter', filter, bool),
            ('storage_domain', storage_domain, types.StorageDomain),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            filter=filter,
            storage_domain=storage_domain,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'move', None, headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a disk.


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def sparsify(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Sparsify the disk.
        Sparsification frees space in the disk image that is not used by its
        filesystem. As a result, the image will occupy less space on the storage.
        Currently sparsification works only on disks without snapshots. Disks
        having derived disks are also not allowed.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'sparsify', None, headers, query, wait)

    def update(
        self,
        disk,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This operation updates the disk with the appropriate parameters.
        The only field that can be updated is `qcow_version`.
        For example, update disk can be facilitated using the following request:
        [source]
        ----
        PUT /ovirt-engine/api/disks/123
        ----
        With a request body like this:
        [source,xml]
        ----
        <disk>
          <qcow_version>qcow2_v3</qcow_version>
        </disk>
        ----
        Since the backend operation is asynchronous the disk element which will be returned
        to the user might not be synced with the changed properties.


        This method supports the following parameters:

        `disk`:: The update to apply to the disk.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('disk', disk, types.Disk),
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_update(disk, headers, query, wait)

    def permissions_service(self):
        """
        Reference to the service that manages the permissions assigned to the disk.

        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'DiskService:%s' % self._path


class EngineKatelloErrataService(KatelloErrataService):
    """
    A service to manage Katello errata assigned to the engine.
    The information is retrieved from Katello.

    """

    def __init__(self, connection, path):
        super(EngineKatelloErrataService, self).__init__(connection, path)
        self._katello_erratum_service = None

    def list(
        self,
        max=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Retrieves the representation of the Katello errata.
        [source]
        ----
        GET /ovirt-engine/api/katelloerrata
        ----
        You will receive response in XML like this one:
        [source,xml]
        ----
        <katello_errata>
          <katello_erratum href="/ovirt-engine/api/katelloerrata/123" id="123">
            <name>RHBA-2013:XYZ</name>
            <description>The description of the erratum</description>
            <title>some bug fix update</title>
            <type>bugfix</type>
            <issued>2013-11-20T02:00:00.000+02:00</issued>
            <solution>Few guidelines regarding the solution</solution>
            <summary>Updated packages that fix one bug are now available for XYZ</summary>
            <packages>
              <package>
                <name>libipa_hbac-1.9.2-82.11.el6_4.i686</name>
              </package>
              ...
            </packages>
          </katello_erratum>
          ...
        </katello_errata>
        ----


        This method supports the following parameters:

        `max`:: Sets the maximum number of errata to return. If not specified all the errata are returned.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('max', max, int),
        ])

        # Build the URL:
        query = query or {}
        if max is not None:
            max = Writer.render_integer(max)
            query['max'] = max

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def katello_erratum_service(self, id):
        """
        Reference to the Katello erratum service.
        Use this service to view the erratum by its id.

        """
        Service._check_types([
            ('id', id, str),
        ])
        return KatelloErratumService(self._connection, '%s/%s' % (self._path, id))

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        index = path.find('/')
        if index == -1:
            return self.katello_erratum_service(path)
        return self.katello_erratum_service(path[:index]).service(path[index + 1:])

    def __str__(self):
        return 'EngineKatelloErrataService:%s' % self._path


class ExternalHostProviderService(ExternalProviderService):
    """
    """

    def __init__(self, connection, path):
        super(ExternalHostProviderService, self).__init__(connection, path)
        self._certificates_service = None
        self._compute_resources_service = None
        self._discovered_hosts_service = None
        self._host_groups_service = None
        self._hosts_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def import_certificates(
        self,
        certificates=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('certificates', certificates, list),
        ])

        # Populate the action:
        action = types.Action(
            certificates=certificates,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'importcertificates', None, headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def test_connectivity(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the test should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'testconnectivity', None, headers, query, wait)

    def update(
        self,
        provider,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
            ('provider', provider, types.ExternalHostProvider),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(provider, headers, query, wait)

    def certificates_service(self):
        """
        """
        return ExternalProviderCertificatesService(self._connection, '%s/certificates' % self._path)

    def compute_resources_service(self):
        """
        """
        return ExternalComputeResourcesService(self._connection, '%s/computeresources' % self._path)

    def discovered_hosts_service(self):
        """
        """
        return ExternalDiscoveredHostsService(self._connection, '%s/discoveredhosts' % self._path)

    def host_groups_service(self):
        """
        """
        return ExternalHostGroupsService(self._connection, '%s/hostgroups' % self._path)

    def hosts_service(self):
        """
        """
        return ExternalHostsService(self._connection, '%s/hosts' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'certificates':
            return self.certificates_service()
        if path.startswith('certificates/'):
            return self.certificates_service().service(path[13:])
        if path == 'computeresources':
            return self.compute_resources_service()
        if path.startswith('computeresources/'):
            return self.compute_resources_service().service(path[17:])
        if path == 'discoveredhosts':
            return self.discovered_hosts_service()
        if path.startswith('discoveredhosts/'):
            return self.discovered_hosts_service().service(path[16:])
        if path == 'hostgroups':
            return self.host_groups_service()
        if path.startswith('hostgroups/'):
            return self.host_groups_service().service(path[11:])
        if path == 'hosts':
            return self.hosts_service()
        if path.startswith('hosts/'):
            return self.hosts_service().service(path[6:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'ExternalHostProviderService:%s' % self._path


class GlusterBrickService(MeasurableService):
    """
    This service manages a single gluster brick.

    """

    def __init__(self, connection, path):
        super(GlusterBrickService, self).__init__(connection, path)
        self._statistics_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get details of a brick.
        Retrieves status details of brick from underlying gluster volume with header `All-Content` set to `true`. This is
        the equivalent of running `gluster volume status <volumename> <brickname> detail`.
        For example, to get the details of brick `234` of gluster volume `123`, send a request like this:
        [source]
        ----
        GET /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks/234
        ----
        Which will return a response body like this:
        [source,xml]
        ----
        <brick id="234">
          <name>host1:/rhgs/data/brick1</name>
          <brick_dir>/rhgs/data/brick1</brick_dir>
          <server_id>111</server_id>
          <status>up</status>
          <device>/dev/mapper/RHGS_vg1-lv_vmaddldisks</device>
          <fs_name>xfs</fs_name>
          <gluster_clients>
            <gluster_client>
              <bytes_read>2818417648</bytes_read>
              <bytes_written>1384694844</bytes_written>
              <client_port>1011</client_port>
              <host_name>client2</host_name>
            </gluster_client>
          </gluster_clients>
          <memory_pools>
            <memory_pool>
              <name>data-server:fd_t</name>
              <alloc_count>1626348</alloc_count>
              <cold_count>1020</cold_count>
              <hot_count>4</hot_count>
              <max_alloc>23</max_alloc>
              <max_stdalloc>0</max_stdalloc>
              <padded_size>140</padded_size>
              <pool_misses>0</pool_misses>
            </memory_pool>
          </memory_pools>
          <mnt_options>rw,seclabel,noatime,nodiratime,attr2,inode64,sunit=512,swidth=2048,noquota</mnt_options>
          <pid>25589</pid>
          <port>49155</port>
        </brick>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes a brick.
        Removes a brick from the underlying gluster volume and deletes entries from database. This can be used only when
        removing a single brick without data migration. To remove multiple bricks and with data migration, use
        <<services/gluster_bricks/methods/migrate, migrate>> instead.
        For example, to delete brick `234` from gluster volume `123`, send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/clusters/567/glustervolumes/123/glusterbricks/234
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def replace(
        self,
        async=None,
        force=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Replaces this brick with a new one.
        IMPORTANT: This operation has been deprecated since version 3.5 of the engine and will be removed in the future.
        Use <<services/gluster_bricks/methods/add, add brick(s)>> and
        <<services/gluster_bricks/methods/migrate, migrate brick(s)>> instead.


        This method supports the following parameters:

        `async`:: Indicates if the replacement should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('force', force, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            force=force,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'replace', None, headers, query, wait)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'GlusterBrickService:%s' % self._path


class GlusterVolumeService(MeasurableService):
    """
    This service manages a single gluster volume.

    """

    def __init__(self, connection, path):
        super(GlusterVolumeService, self).__init__(connection, path)
        self._gluster_bricks_service = None
        self._statistics_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get the gluster volume details.
        For example, to get details of a gluster volume with identifier `123` in cluster `456`, send a request like this:
        [source]
        ----
        GET /ovirt-engine/api/clusters/456/glustervolumes/123
        ----
        This GET request will return the following output:
        [source,xml]
        ----
        <gluster_volume id="123">
         <name>data</name>
         <link href="/ovirt-engine/api/clusters/456/glustervolumes/123/glusterbricks" rel="glusterbricks"/>
         <disperse_count>0</disperse_count>
         <options>
           <option>
             <name>storage.owner-gid</name>
             <value>36</value>
           </option>
           <option>
             <name>performance.io-cache</name>
             <value>off</value>
           </option>
           <option>
             <name>cluster.data-self-heal-algorithm</name>
             <value>full</value>
           </option>
         </options>
         <redundancy_count>0</redundancy_count>
         <replica_count>3</replica_count>
         <status>up</status>
         <stripe_count>0</stripe_count>
         <transport_types>
           <transport_type>tcp</transport_type>
         </transport_types>
         <volume_type>replicate</volume_type>
         </gluster_volume>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def get_profile_statistics(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get gluster volume profile statistics.
        For example, to get profile statistics for a gluster volume with identifier `123` in cluster `456`, send a
        request like this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/456/glustervolumes/123/getprofilestatistics
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'getprofilestatistics', 'details', headers, query, wait)

    def rebalance(
        self,
        async=None,
        fix_layout=None,
        force=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Rebalance the gluster volume.
        Rebalancing a gluster volume helps to distribute the data evenly across all the bricks. After expanding or
        shrinking a gluster volume (without migrating data), we need to rebalance the data among the bricks. In a
        non-replicated volume, all bricks should be online to perform the rebalance operation. In a replicated volume, at
        least one of the bricks in the replica should be online.
        For example, to rebalance a gluster volume with identifier `123` in cluster `456`, send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/456/glustervolumes/123/rebalance
        ----


        This method supports the following parameters:

        `fix_layout`:: If set to true, rebalance will only fix the layout so that new data added to the volume is distributed
        across all the hosts. But it will not migrate/rebalance the existing data. Default is `false`.

        `force`:: Indicates if the rebalance should be force started. The rebalance command can be executed with the force
        option even when the older clients are connected to the cluster. However, this could lead to a data loss
        situation. Default is `false`.

        `async`:: Indicates if the rebalance should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('fix_layout', fix_layout, bool),
            ('force', force, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            fix_layout=fix_layout,
            force=force,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'rebalance', None, headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Removes the gluster volume.
        For example, to remove a volume with identifier `123` in cluster `456`, send a request like this:
        [source]
        ----
        DELETE /ovirt-engine/api/clusters/456/glustervolumes/123
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def reset_all_options(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Resets all the options set in the gluster volume.
        For example, to reset all options in a gluster volume with identifier `123` in cluster `456`, send a request like
        this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/456/glustervolumes/123/resetalloptions
        ----


        This method supports the following parameters:

        `async`:: Indicates if the reset should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'resetalloptions', None, headers, query, wait)

    def reset_option(
        self,
        async=None,
        force=None,
        option=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Resets a particular option in the gluster volume.
        For example, to reset a particular option `option1` in a gluster volume with identifier `123` in cluster `456`,
        send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/456/glustervolumes/123/resetoption
        ----
        With the following request body:
        [source,xml]
        ----
        <action>
         <option name="option1"/>
        </action>
        ----


        This method supports the following parameters:

        `option`:: Option to reset.

        `async`:: Indicates if the reset should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('force', force, bool),
            ('option', option, types.Option),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            force=force,
            option=option,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'resetoption', None, headers, query, wait)

    def set_option(
        self,
        async=None,
        option=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Sets a particular option in the gluster volume.
        For example, to set `option1` with value `value1` in a gluster volume with identifier `123` in cluster `456`,
        send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/456/glustervolumes/123/setoption
        ----
        With the following request body:
        [source,xml]
        ----
        <action>
         <option name="option1" value="value1"/>
        </action>
        ----


        This method supports the following parameters:

        `option`:: Option to set.

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('option', option, types.Option),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            option=option,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'setoption', None, headers, query, wait)

    def start(
        self,
        async=None,
        force=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Starts the gluster volume.
        A Gluster Volume should be started to read/write data. For example, to start a gluster volume with identifier
        `123` in cluster `456`, send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/456/glustervolumes/123/start
        ----


        This method supports the following parameters:

        `force`:: Indicates if the volume should be force started. If a gluster volume is started already but few/all bricks
        are down then force start can be used to bring all the bricks up. Default is `false`.

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('force', force, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            force=force,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'start', None, headers, query, wait)

    def start_profile(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Start profiling the gluster volume.
        For example, to start profiling a gluster volume with identifier `123` in cluster `456`, send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/456/glustervolumes/123/startprofile
        ----


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'startprofile', None, headers, query, wait)

    def stop(
        self,
        async=None,
        force=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Stops the gluster volume.
        Stopping a volume will make its data inaccessible.
        For example, to stop a gluster volume with identifier `123` in cluster `456`, send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/456/glustervolumes/123/stop
        ----


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('force', force, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            force=force,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'stop', None, headers, query, wait)

    def stop_profile(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Stop profiling the gluster volume.
        For example, to stop profiling a gluster volume with identifier `123` in cluster `456`, send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/456/glustervolumes/123/stopprofile
        ----


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'stopprofile', None, headers, query, wait)

    def stop_rebalance(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Stop rebalancing the gluster volume.
        For example, to stop rebalancing a gluster volume with identifier `123` in cluster `456`, send a request like
        this:
        [source]
        ----
        POST /ovirt-engine/api/clusters/456/glustervolumes/123/stoprebalance
        ----


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'stoprebalance', None, headers, query, wait)

    def gluster_bricks_service(self):
        """
        Reference to a service managing gluster bricks.

        """
        return GlusterBricksService(self._connection, '%s/glusterbricks' % self._path)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'glusterbricks':
            return self.gluster_bricks_service()
        if path.startswith('glusterbricks/'):
            return self.gluster_bricks_service().service(path[14:])
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'GlusterVolumeService:%s' % self._path


class HostService(MeasurableService):
    """
    A service to manage a host.

    """

    def __init__(self, connection, path):
        super(HostService, self).__init__(connection, path)
        self._affinity_labels_service = None
        self._devices_service = None
        self._fence_agents_service = None
        self._hooks_service = None
        self._katello_errata_service = None
        self._network_attachments_service = None
        self._nics_service = None
        self._numa_nodes_service = None
        self._permissions_service = None
        self._statistics_service = None
        self._storage_service = None
        self._storage_connection_extensions_service = None
        self._tags_service = None
        self._unmanaged_networks_service = None

    def activate(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Activate the host for use, such as running virtual machines.


        This method supports the following parameters:

        `async`:: Indicates if the activation should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'activate', None, headers, query, wait)

    def approve(
        self,
        async=None,
        cluster=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Approve a pre-installed Hypervisor host for usage in the virtualization environment.
        This action also accepts an optional cluster element to define the target cluster for this host.


        This method supports the following parameters:

        `async`:: Indicates if the approval should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('cluster', cluster, types.Cluster),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            cluster=cluster,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'approve', None, headers, query, wait)

    def commit_net_config(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Marks the network configuration as good and persists it inside the host.
        An API user commits the network configuration to persist a host network interface attachment or detachment, or
        persist the creation and deletion of a bonded interface.
        IMPORTANT: Networking configuration is only committed after the engine has established that host connectivity is
        not lost as a result of the configuration changes. If host connectivity is lost, the host requires a reboot and
        automatically reverts to the previous networking configuration.
        For example, to commit the network configuration of host with id `123` send a request like this:
        [source]
        ----
        POST /ovirt-engine/api/hosts/123/commitnetconfig
        ----
        With a request body like this:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'commitnetconfig', None, headers, query, wait)

    def deactivate(
        self,
        async=None,
        reason=None,
        stop_gluster_service=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Deactivate the host to perform maintenance tasks.


        This method supports the following parameters:

        `async`:: Indicates if the deactivation should be performed asynchronously.

        `stop_gluster_service`:: Indicates if the gluster service should be stopped as part of deactivating the host. It can be used while
        performing maintenance operations on the gluster host. Default value for this variable is `false`.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('reason', reason, str),
            ('stop_gluster_service', stop_gluster_service, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            reason=reason,
            stop_gluster_service=stop_gluster_service,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'deactivate', None, headers, query, wait)

    def enroll_certificate(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Enroll certificate of the host. Useful in case you get a warning that it is about to, or already expired.


        This method supports the following parameters:

        `async`:: Indicates if the enrollment should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'enrollcertificate', None, headers, query, wait)

    def fence(
        self,
        async=None,
        fence_type=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Controls host's power management device.
        For example, let's assume you want to start the host. This can be done via:
        [source]
        ----
        #!/bin/sh -ex
        url="https://engine.example.com/ovirt-engine/api"
        user="admin@internal"
        password="..."
        curl \
        --verbose \
        --cacert /etc/pki/ovirt-engine/ca.pem \
        --user "${user}:${password}" \
        --request POST \
        --header "Version: 4" \
        --header "Content-Type: application/xml" \
        --header "Accept: application/xml" \
        --data '
        <action>
          <fence_type>start</fence_type>
        </action>
        ' \
        "${url}/hosts/123/fence"
        ----


        This method supports the following parameters:

        `async`:: Indicates if the fencing should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('fence_type', fence_type, str),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            fence_type=fence_type,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'fence', 'power_management', headers, query, wait)

    def force_select_spm(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Manually set a host as the storage pool manager (SPM).
        [source]
        ----
        POST /ovirt-engine/api/hosts/123/forceselectspm
        ----
        With a request body like this:
        [source,xml]
        ----
        <action/>
        ----


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'forceselectspm', None, headers, query, wait)

    def get(
        self,
        filter=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Get the host details.


        This method supports the following parameters:

        `filter`:: Indicates if the results should be filtered according to the permissions of the user.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('filter', filter, bool),
        ])

        # Build the URL:
        query = query or {}
        if filter is not None:
            filter = Writer.render_boolean(filter)
            query['filter'] = filter

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def install(
        self,
        async=None,
        deploy_hosted_engine=None,
        host=None,
        image=None,
        root_password=None,
        ssh=None,
        undeploy_hosted_engine=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Install VDSM and related software on the host. The host type defines additional parameters for the action.
        Example of installing a host, using `curl` and JSON, plain:
        [source,bash]
        ----
        curl \
        --verbose \
        --cacert /etc/pki/ovirt-engine/ca.pem \
        --request PUT \
        --header "Content-Type: application/json" \
        --header "Accept: application/json" \
        --header "Version: 4" \
        --user "admin@internal:..." \
        --data '
        {
          "root_password": "myrootpassword"
        }
        ' \
        "https://engine.example.com/ovirt-engine/api/hosts/123"
        ----
        Example of installing a host, using `curl` and JSON, with hosted engine components:
        [source,bash]
        ----
        curl \
        curl \
        --verbose \
        --cacert /etc/pki/ovirt-engine/ca.pem \
        --request PUT \
        --header "Content-Type: application/json" \
        --header "Accept: application/json" \
        --header "Version: 4" \
        --user "admin@internal:..." \
        --data '
        {
          "root_password": "myrootpassword"
        }
        ' \
        "https://engine.example.com/ovirt-engine/api/hosts/123?deploy_hosted_engine=true"
        ----


        This method supports the following parameters:

        `root_password`:: The password of of the `root` user, used to connect to the host via SSH.

        `ssh`:: The SSH details used to connect to the host.

        `host`:: This `override_iptables` property is used to indicate if the firewall configuration should be
        replaced by the default one.

        `image`:: When installing an oVirt node a image ISO file is needed.

        `async`:: Indicates if the installation should be performed asynchronously.

        `deploy_hosted_engine`:: When set to `true` it means this host should deploy also hosted
        engine components. Missing value is treated as `true` i.e deploy.
        Omitting this parameter means `false` and will perform no operation
        in hosted engine area.

        `undeploy_hosted_engine`:: When set to `true` it means this host should un-deploy hosted engine
        components and this host will not function as part of the High
        Availability cluster. Missing value is treated as `true` i.e un-deploy
        Omitting this parameter means `false` and will perform no operation
        in hosted engine area.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('deploy_hosted_engine', deploy_hosted_engine, bool),
            ('host', host, types.Host),
            ('image', image, str),
            ('root_password', root_password, str),
            ('ssh', ssh, types.Ssh),
            ('undeploy_hosted_engine', undeploy_hosted_engine, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            deploy_hosted_engine=deploy_hosted_engine,
            host=host,
            image=image,
            root_password=root_password,
            ssh=ssh,
            undeploy_hosted_engine=undeploy_hosted_engine,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'install', None, headers, query, wait)

    def iscsi_discover(
        self,
        async=None,
        iscsi=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Discover iSCSI targets on the host, using the initiator details.


        This method supports the following parameters:

        `iscsi`:: The target iSCSI device.

        `async`:: Indicates if the discovery should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('iscsi', iscsi, types.IscsiDetails),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            iscsi=iscsi,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'iscsidiscover', 'iscsi_targets', headers, query, wait)

    def iscsi_login(
        self,
        async=None,
        iscsi=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Login to iSCSI targets on the host, using the target details.


        This method supports the following parameters:

        `iscsi`:: The target iSCSI device.

        `async`:: Indicates if the login should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('iscsi', iscsi, types.IscsiDetails),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            iscsi=iscsi,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'iscsilogin', None, headers, query, wait)

    def refresh(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Refresh the host devices and capabilities.


        This method supports the following parameters:

        `async`:: Indicates if the refresh should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'refresh', None, headers, query, wait)

    def remove(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Remove the host from the system.
        [source]
        ----
        #!/bin/sh -ex
        url="https://engine.example.com/ovirt-engine/api"
        user="admin@internal"
        password="..."
        curl \
        --verbose \
        --cacert /etc/pki/ovirt-engine/ca.pem \
        --user "${user}:${password}" \
        --request DELETE \
        --header "Version: 4" \
        "${url}/hosts/1ff7a191-2f3b-4eff-812b-9f91a30c3acc"
        ----


        This method supports the following parameters:

        `async`:: Indicates if the remove should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        self._internal_remove(headers, query, wait)

    def setup_networks(
        self,
        async=None,
        check_connectivity=None,
        connectivity_timeout=None,
        modified_bonds=None,
        modified_labels=None,
        modified_network_attachments=None,
        removed_bonds=None,
        removed_labels=None,
        removed_network_attachments=None,
        synchronized_network_attachments=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method is used to change the configuration of the network interfaces of a host.
        For example, lets assume that you have a host with three network interfaces `eth0`, `eth1` and `eth2` and that
        you want to configure a new bond using `eth0` and `eth1`, and put a VLAN on top of it. Using a simple shell
        script and the `curl` command line HTTP client that can be done as follows:
        [source]
        ----
        #!/bin/sh -ex
        url="https://engine.example.com/ovirt-engine/api"
        user="admin@internal"
        password="..."
        curl \
        --verbose \
        --cacert /etc/pki/ovirt-engine/ca.pem \
        --user "${user}:${password}" \
        --request POST \
        --header "Version: 4" \
        --header "Content-Type: application/xml" \
        --header "Accept: application/xml" \
        --data '
        <action>
          <modified_bonds>
            <host_nic>
              <name>bond0</name>
              <bonding>
                <options>
                  <option>
                    <name>mode</name>
                    <value>4</value>
                  </option>
                  <option>
                    <name>miimon</name>
                    <value>100</value>
                  </option>
                </options>
                <slaves>
                  <host_nic>
                    <name>eth1</name>
                  </host_nic>
                  <host_nic>
                    <name>eth2</name>
                  </host_nic>
                </slaves>
              </bonding>
            </host_nic>
          </modified_bonds>
          <modified_network_attachments>
            <network_attachment>
              <network>
                <name>myvlan</name>
              </network>
              <host_nic>
                <name>bond0</name>
              </host_nic>
              <ip_address_assignments>
                <assignment_method>static</assignment_method>
                <ip_address_assignment>
                  <ip>
                    <address>192.168.122.10</address>
                    <netmask>255.255.255.0</netmask>
                  </ip>
                </ip_address_assignment>
              </ip_address_assignments>
              <dns_resolver_configuration>
                <name_servers>
                  <name_server>1.1.1.1</name_server>
                  <name_server>2.2.2.2</name_server>
                </name_servers>
              </dns_resolver_configuration>
            </network_attachment>
          </modified_network_attachments>
         </action>
        ' \
        "${url}/hosts/1ff7a191-2f3b-4eff-812b-9f91a30c3acc/setupnetworks"
        ----
        Note that this is valid for version 4 of the API. In previous versions some elements were represented as XML
        attributes instead of XML elements. In particular the `options` and `ip` elements were represented as follows:
        [source,xml]
        ----
        <options name="mode" value="4"/>
        <options name="miimon" value="100"/>
        <ip address="192.168.122.10" netmask="255.255.255.0"/>
        ----
        Using the Python SDK the same can be done with the following code:
        [source,python]
        ----
        # Find the service that manages the collection of hosts:
        hosts_service = connection.system_service().hosts_service()
        # Find the host:
        host = hosts_service.list(search='name=myhost')[0]
        # Find the service that manages the host:
        host_service = hosts_service.host_service(host.id)
        # Configure the network adding a bond with two slaves and attaching it to a
        # network with an static IP address:
        host_service.setup_networks(
            modified_bonds=[
                types.HostNic(
                    name='bond0',
                    bonding=types.Bonding(
                        options=[
                            types.Option(
                                name='mode',
                                value='4',
                            ),
                            types.Option(
                                name='miimon',
                                value='100',
                            ),
                        ],
                        slaves=[
                            types.HostNic(
                                name='eth1',
                            ),
                            types.HostNic(
                                name='eth2',
                            ),
                        ],
                    ),
                ),
            ],
            modified_network_attachments=[
                types.NetworkAttachment(
                    network=types.Network(
                        name='myvlan',
                    ),
                    host_nic=types.HostNic(
                        name='bond0',
                    ),
                    ip_address_assignments=[
                        types.IpAddressAssignment(
                            assignment_method=types.BootProtocol.STATIC,
                            ip=types.Ip(
                                address='192.168.122.10',
                                netmask='255.255.255.0',
                            ),
                        ),
                    ],
                    dns_resolver_configuration=types.DnsResolverConfiguration(
                        name_servers=[
                            '1.1.1.1',
                            '2.2.2.2',
                        ],
                    ),
                ),
            ],
        )
        # After modifying the network configuration it is very important to make it
        # persistent:
        host_service.commit_net_config()
        ----
        IMPORTANT: To make sure that the network configuration has been saved in the host, and that it will be applied
        when the host is rebooted, remember to call <<services/host/methods/commit_net_config, commitnetconfig>>.


        This method supports the following parameters:

        `async`:: Indicates if the action should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('check_connectivity', check_connectivity, bool),
            ('connectivity_timeout', connectivity_timeout, int),
            ('modified_bonds', modified_bonds, list),
            ('modified_labels', modified_labels, list),
            ('modified_network_attachments', modified_network_attachments, list),
            ('removed_bonds', removed_bonds, list),
            ('removed_labels', removed_labels, list),
            ('removed_network_attachments', removed_network_attachments, list),
            ('synchronized_network_attachments', synchronized_network_attachments, list),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            check_connectivity=check_connectivity,
            connectivity_timeout=connectivity_timeout,
            modified_bonds=modified_bonds,
            modified_labels=modified_labels,
            modified_network_attachments=modified_network_attachments,
            removed_bonds=removed_bonds,
            removed_labels=removed_labels,
            removed_network_attachments=removed_network_attachments,
            synchronized_network_attachments=synchronized_network_attachments,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'setupnetworks', None, headers, query, wait)

    def unregistered_storage_domains_discover(
        self,
        async=None,
        iscsi=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        This method supports the following parameters:

        `async`:: Indicates if the discovery should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('iscsi', iscsi, types.IscsiDetails),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            iscsi=iscsi,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'unregisteredstoragedomainsdiscover', 'storage_domains', headers, query, wait)

    def update(
        self,
        host,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Update the host properties.
        For example, to update a the kernel command line of a host send a request like this:
        [source]
        ----
        PUT /ovirt-engine/api/hosts/123
        ----
        With request body like this:
        [source, xml]
        ----
        <host>
          <os>
            <custom_kernel_cmdline>vfio_iommu_type1.allow_unsafe_interrupts=1</custom_kernel_cmdline>
          </os>
        </host>
        ----


        """
        # Check the types of the parameters:
        Service._check_types([
            ('host', host, types.Host),
            ('async', async, bool),
        ])

        # Build the URL:
        query = query or {}
        if async is not None:
            async = Writer.render_boolean(async)
            query['async'] = async

        # Send the request and wait for the response:
        return self._internal_update(host, headers, query, wait)

    def upgrade(
        self,
        async=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Upgrade VDSM and selected software on the host.


        This method supports the following parameters:

        `async`:: Indicates if the upgrade should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'upgrade', None, headers, query, wait)

    def upgrade_check(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        Check if there are upgrades available for the host. If there are upgrades
        available an icon will be displayed next to host status icon in the webadmin.
        Audit log messages are also added to indicate the availability of upgrades.
        The upgrade can be started from the webadmin or by using the
        <<services/host/methods/upgrade, upgrade>> host action.


        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Populate the action:
        action = types.Action(
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'upgradecheck', None, headers, query, wait)

    def affinity_labels_service(self):
        """
        List of scheduling labels assigned to this host.

        """
        return AssignedAffinityLabelsService(self._connection, '%s/affinitylabels' % self._path)

    def devices_service(self):
        """
        Reference to the host devices service.
        Use this service to view the devices of the host object.

        """
        return HostDevicesService(self._connection, '%s/devices' % self._path)

    def fence_agents_service(self):
        """
        Reference to the fence agents service.
        Use this service to manage fence and power management agents on the host object.

        """
        return FenceAgentsService(self._connection, '%s/fenceagents' % self._path)

    def hooks_service(self):
        """
        Reference to the host hooks service.
        Use this service to view the hooks available in the host object.

        """
        return HostHooksService(self._connection, '%s/hooks' % self._path)

    def katello_errata_service(self):
        """
        Reference to the service that can show the applicable errata available on the host.
        This information is taken from Katello.

        """
        return KatelloErrataService(self._connection, '%s/katelloerrata' % self._path)

    def network_attachments_service(self):
        """
        Reference to the network attachments service. You can use this service to attach
        Logical networks to host interfaces.

        """
        return NetworkAttachmentsService(self._connection, '%s/networkattachments' % self._path)

    def nics_service(self):
        """
        Reference to the service that manages the network interface devices on the host.

        """
        return HostNicsService(self._connection, '%s/nics' % self._path)

    def numa_nodes_service(self):
        """
        Reference to the service that manage NUMA nodes for the host.

        """
        return HostNumaNodesService(self._connection, '%s/numanodes' % self._path)

    def permissions_service(self):
        """
        Reference to the host permission service.
        Use this service to manage permissions on the host object.

        """
        return AssignedPermissionsService(self._connection, '%s/permissions' % self._path)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def storage_service(self):
        """
        Reference to the service that manage hosts storage.

        """
        return HostStorageService(self._connection, '%s/storage' % self._path)

    def storage_connection_extensions_service(self):
        """
        Reference to storage connection extensions.

        """
        return StorageServerConnectionExtensionsService(self._connection, '%s/storageconnectionextensions' % self._path)

    def tags_service(self):
        """
        Reference to the host tags service.
        Use this service to manage tags on the host object.

        """
        return AssignedTagsService(self._connection, '%s/tags' % self._path)

    def unmanaged_networks_service(self):
        """
        Reference to unmanaged networks.

        """
        return UnmanagedNetworksService(self._connection, '%s/unmanagednetworks' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'affinitylabels':
            return self.affinity_labels_service()
        if path.startswith('affinitylabels/'):
            return self.affinity_labels_service().service(path[15:])
        if path == 'devices':
            return self.devices_service()
        if path.startswith('devices/'):
            return self.devices_service().service(path[8:])
        if path == 'fenceagents':
            return self.fence_agents_service()
        if path.startswith('fenceagents/'):
            return self.fence_agents_service().service(path[12:])
        if path == 'hooks':
            return self.hooks_service()
        if path.startswith('hooks/'):
            return self.hooks_service().service(path[6:])
        if path == 'katelloerrata':
            return self.katello_errata_service()
        if path.startswith('katelloerrata/'):
            return self.katello_errata_service().service(path[14:])
        if path == 'networkattachments':
            return self.network_attachments_service()
        if path.startswith('networkattachments/'):
            return self.network_attachments_service().service(path[19:])
        if path == 'nics':
            return self.nics_service()
        if path.startswith('nics/'):
            return self.nics_service().service(path[5:])
        if path == 'numanodes':
            return self.numa_nodes_service()
        if path.startswith('numanodes/'):
            return self.numa_nodes_service().service(path[10:])
        if path == 'permissions':
            return self.permissions_service()
        if path.startswith('permissions/'):
            return self.permissions_service().service(path[12:])
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        if path == 'storage':
            return self.storage_service()
        if path.startswith('storage/'):
            return self.storage_service().service(path[8:])
        if path == 'storageconnectionextensions':
            return self.storage_connection_extensions_service()
        if path.startswith('storageconnectionextensions/'):
            return self.storage_connection_extensions_service().service(path[28:])
        if path == 'tags':
            return self.tags_service()
        if path.startswith('tags/'):
            return self.tags_service().service(path[5:])
        if path == 'unmanagednetworks':
            return self.unmanaged_networks_service()
        if path.startswith('unmanagednetworks/'):
            return self.unmanaged_networks_service().service(path[18:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'HostService:%s' % self._path


class HostNicService(MeasurableService):
    """
    A service to manage a network interface of a host.

    """

    def __init__(self, connection, path):
        super(HostNicService, self).__init__(connection, path)
        self._network_attachments_service = None
        self._network_labels_service = None
        self._statistics_service = None
        self._virtual_function_allowed_labels_service = None
        self._virtual_function_allowed_networks_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def update_virtual_functions_configuration(
        self,
        async=None,
        virtual_functions_configuration=None,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        The action updates virtual function configuration in case the current resource represents an SR-IOV enabled NIC.
        The input should be consisted of at least one of the following properties:
        - `allNetworksAllowed`
        - `numberOfVirtualFunctions`
        Please see the `HostNicVirtualFunctionsConfiguration` type for the meaning of the properties.


        This method supports the following parameters:

        `async`:: Indicates if the update should be performed asynchronously.

        `headers`:: Additional HTTP headers.

        `query`:: Additional URL query parameters.

        `wait`:: If `True` wait for the response.
        """
        # Check the types of the parameters:
        Service._check_types([
            ('async', async, bool),
            ('virtual_functions_configuration', virtual_functions_configuration, types.HostNicVirtualFunctionsConfiguration),
        ])

        # Populate the action:
        action = types.Action(
            async=async,
            virtual_functions_configuration=virtual_functions_configuration,
        )

        # Send the request and wait for the response:
        return self._internal_action(action, 'updatevirtualfunctionsconfiguration', None, headers, query, wait)

    def network_attachments_service(self):
        """
        Reference to the service that manages the network attachments assigned to this network interface.

        """
        return NetworkAttachmentsService(self._connection, '%s/networkattachments' % self._path)

    def network_labels_service(self):
        """
        Reference to the service that manages the network labels assigned to this network interface.

        """
        return NetworkLabelsService(self._connection, '%s/networklabels' % self._path)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def virtual_function_allowed_labels_service(self):
        """
        Retrieves sub-collection resource of network labels that are allowed on an the virtual functions
        in case that the current resource represents an SR-IOV physical function NIC.

        """
        return NetworkLabelsService(self._connection, '%s/virtualfunctionallowedlabels' % self._path)

    def virtual_function_allowed_networks_service(self):
        """
        Retrieves sub-collection resource of networks that are allowed on an the virtual functions
        in case that the current resource represents an SR-IOV physical function NIC.

        """
        return VirtualFunctionAllowedNetworksService(self._connection, '%s/virtualfunctionallowednetworks' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'networkattachments':
            return self.network_attachments_service()
        if path.startswith('networkattachments/'):
            return self.network_attachments_service().service(path[19:])
        if path == 'networklabels':
            return self.network_labels_service()
        if path.startswith('networklabels/'):
            return self.network_labels_service().service(path[14:])
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        if path == 'virtualfunctionallowedlabels':
            return self.virtual_function_allowed_labels_service()
        if path.startswith('virtualfunctionallowedlabels/'):
            return self.virtual_function_allowed_labels_service().service(path[29:])
        if path == 'virtualfunctionallowednetworks':
            return self.virtual_function_allowed_networks_service()
        if path.startswith('virtualfunctionallowednetworks/'):
            return self.virtual_function_allowed_networks_service().service(path[31:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'HostNicService:%s' % self._path


class HostNumaNodeService(MeasurableService):
    """
    """

    def __init__(self, connection, path):
        super(HostNumaNodeService, self).__init__(connection, path)
        self._statistics_service = None

    def get(
        self,
        headers=None,
        query=None,
        wait=True,
    ):
        """
        """
        # Check the types of the parameters:
        Service._check_types([
        ])

        # Build the URL:
        query = query or {}

        # Send the request and wait for the response:
        return self._internal_get(headers, query, wait)

    def statistics_service(self):
        """
        """
        return StatisticsService(self._connection, '%s/statistics' % self._path)

    def service(self, path):
        """
        Service locator method, returns individual service on which the URI is dispatched.
        """
        if not path:
            return self
        if path == 'statistics':
            return self.statistics_service()
        if path.startswith('statistics/'):
            return self.statistics_service().service(path[11:])
        raise Error('The path \"%s\" doesn\'t correspond to any service' % path)

    def __str__(self):
        return 'HostNumaNodeService:%s' % self._path
