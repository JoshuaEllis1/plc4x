using System;
using System.Net.Sockets;
using System.Net;
//using System.IO.Ports;
using System.Reflection;
using System.Text;
using System.Collections.Generic;
using org.apache.plc4net.exceptions;

namespace org.apache.plc4net.drivers.modbus
{
    public class ModbusDriver
    {
        /// <summary>
        /// 
        /// </summary>
        public void TCPConnect()
        {
            TcpClient tcpClient = new TcpClient();
            var connectionClient = tcpClient.BeginConnect("localhost", 502, null, null);
            var successfulConnection = connectionClient.AsyncWaitHandle.WaitOne(1000);
            if (!successfulConnection)
            {
                throw new PlcConnectionException();
            }
            tcpClient.EndConnect(connectionClient);
        
        }

        /// <summary>
        /// 
        /// </summary>
        public void UDPConnect()
        {
            UdpClient udpClient = new UdpClient();
            
        }


        /// <summary>
        /// 
        /// </summary>
        /// <returns></returns>
        public bool IsAvailable()
        {
            System.Net.NetworkInformation.Ping ping = new System.Net.NetworkInformation.Ping();
            IPAddress iPAddress = System.Net.IPAddress.Parse("");
            string bufferData = "abcdefghijklmnopqrstuvwxyzabcdef";
            byte[] bufferFinal = System.Text.Encoding.ASCII.GetBytes(bufferData);
            System.Net.NetworkInformation.PingReply reply = ping.Send(iPAddress,10,bufferFinal);

            if(reply.Status == System.Net.NetworkInformation.IPStatus.Success)
            {
                return true;
            }
            else
            {
                return false;
            }
        }
       
    }
}
