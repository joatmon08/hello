control 'check application binary' do
  impact 1.0
  title 'confirm binary exists'
  desc 'confirm binary exists'
  describe command('which main') do
    its('stdout') { should include '/usr/local/bin/main' }
  end
end

control 'check application security' do
  impact 1.0
  title 'confirm application security user'
  desc 'confirm application security user'
  describe processes('main') do
    it { should exist }
    its('users') { should_not eq ['root'] }
  end
end